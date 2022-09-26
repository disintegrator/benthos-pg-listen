package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func main() {
	if err := mainErr(); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatal(err)
	}
}

func mainErr() error {
	dsn := flag.String("dsn", "", "The connection URL to the PostgreSQL server (e.g. postgres://user:password@localhost:5432/demodb).")
	channel := flag.String("channel", "", "The channel name to listen on.")

	flag.Parse()

	if *dsn == "" {
		return errors.New("-dsn <url> is required")
	}
	if *channel == "" {
		return errors.New("-channel <name> is required")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, sigCancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	defer sigCancel()

	conn, err := pgx.Connect(ctx, *dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	_, err = conn.Exec(ctx, fmt.Sprintf("listen %s", *channel))
	if err != nil {
		return fmt.Errorf("failed to execute LISTEN: %w", err)
	}

	log.Println("Waiting for notifications...")
	for {
		notification, err := conn.WaitForNotification(ctx)
		if err != nil {
			return fmt.Errorf("failed to wait for notification: %w", err)
		}

		err = doSomethingWithNotification(notification)
		if err != nil {
			return fmt.Errorf("failed to process notification: %w", err)
		}
	}
}

func doSomethingWithNotification(n *pgconn.Notification) error {
	fmt.Printf("[%s @ %d]: %s\n", n.Channel, n.PID, n.Payload)
	return nil
}
