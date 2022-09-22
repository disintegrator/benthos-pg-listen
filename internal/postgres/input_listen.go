package postgres

import (
	"context"
	"fmt"

	"github.com/benthosdev/benthos/v4/public/service"
	"github.com/jackc/pgx/v5"
)

type pgListenConfig struct {
	dsn     string
	channel string
}

func newPGListenConfigSpec() *service.ConfigSpec {
	return service.NewConfigSpec().
		Summary("Sets up a listen to a PostgreSQL channel using LISTEN.").
		Field(service.NewStringField("dsn").Description("The connection URL to the PostgreSQL server.")).
		Field(service.NewStringField("channel").Description("The channel name to listen on."))
}

type pgListenInput struct {
	config *pgListenConfig
	conn   *pgx.Conn
}

func newPGListenInputFromConfig(conf *service.ParsedConfig) (*pgListenInput, error) {
	var config pgListenConfig
	var err error

	if config.dsn, err = conf.FieldString("dsn"); err != nil {
		return nil, err
	}
	if config.channel, err = conf.FieldString("channel"); err != nil {
		return nil, err
	}

	return &pgListenInput{config: &config}, nil
}

func (input *pgListenInput) Connect(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, input.config.dsn)
	if err != nil {
		return err
	}

	input.conn = conn

	_, err = conn.Exec(context.Background(), fmt.Sprintf("listen %s", input.config.channel))
	if err != nil {
		return err
	}

	return nil
}

var noopAck = service.AckFunc(func(ctx context.Context, err error) error {
	return nil
})

func (input *pgListenInput) Read(ctx context.Context) (*service.Message, service.AckFunc, error) {
	notification, err := input.conn.WaitForNotification(ctx)
	if err != nil {
		return nil, nil, err
	}

	msg := service.NewMessage([]byte(notification.Payload))
	msg.MetaSet("pg_listen_channel", notification.Channel)
	msg.MetaSet("pg_listen_pid", fmt.Sprint(notification.PID))

	return msg, noopAck, nil
}

func (input *pgListenInput) Close(ctx context.Context) error {
	return input.conn.Close(ctx)
}

func init() {
	service.RegisterInput("postgres_listen", newPGListenConfigSpec(), func(conf *service.ParsedConfig, mgr *service.Resources) (service.Input, error) {
		inp, err := newPGListenInputFromConfig(conf)
		if err != nil {
			return nil, err
		}

		return service.AutoRetryNacks(inp), nil
	})
}
