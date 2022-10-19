package postgres

import (
	"context"

	"github.com/benthosdev/benthos/v4/public/service"
	"github.com/jackc/pgx/v5"
)

type pgNotifyConfig struct {
	dsn         string
	channel     string
	maxInFlight int
}

func newPGNotifyConfigSpec() *service.ConfigSpec {
	return service.NewConfigSpec().
		Summary("Publish a message on a PostgreSQL channel using NOTIFY.").
		Field(service.NewStringField("dsn").Description("The connection URL to the PostgreSQL server.")).
		Field(service.NewStringField("channel").Description("The channel name to notify.")).
		Field(service.NewIntField("max_in_flight").
			Description("The maximum number of message batches to have in flight at a given time.").
			Default(64))
}

type pgNotifyOutput struct {
	config *pgNotifyConfig
	conn   *pgx.Conn
}

func newPGNotifyOutputFromConfig(conf *service.ParsedConfig) (*pgNotifyOutput, error) {
	var config pgNotifyConfig
	var err error

	if config.dsn, err = conf.FieldString("dsn"); err != nil {
		return nil, err
	}
	if config.channel, err = conf.FieldString("channel"); err != nil {
		return nil, err
	}
	if config.maxInFlight, err = conf.FieldInt("max_in_flight"); err != nil {
		return nil, err
	}

	return &pgNotifyOutput{config: &config}, nil
}

func (output *pgNotifyOutput) Connect(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, output.config.dsn)
	if err != nil {
		return err
	}

	output.conn = conn

	return nil
}

func (output *pgNotifyOutput) Write(ctx context.Context, msg *service.Message) error {
	bs, err := msg.AsBytes()
	if err != nil {
		return err
	}

	_, err = output.conn.Exec(ctx, "select pg_notify($1, $2)", output.config.channel, bs)
	return err
}

func (output *pgNotifyOutput) Close(ctx context.Context) error {
	return output.conn.Close(ctx)
}

func init() {
	service.RegisterOutput("postgres_notify", newPGNotifyConfigSpec(), func(conf *service.ParsedConfig, mgr *service.Resources) (service.Output, int, error) {
		out, err := newPGNotifyOutputFromConfig(conf)
		return out, out.config.maxInFlight, err
	})
}
