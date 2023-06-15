package natsBroker

import (
	"context"
	nats "github.com/nats-io/nats.go"
	"surge/config"
	"time"
)

var (
	nc *nats.EncodedConn
)

func Connect(conf config.Config) (err error) {
	var conn *nats.Conn
	opts := nats.Options{
		Secure:         false,
		Servers:        config.GetCnf().Nats,
		PingInterval:   time.Second * 60,
		AllowReconnect: true,
		MaxReconnect:   10,
		ReconnectWait:  5 * time.Second,
		Timeout:        3 * time.Second,
	}

	conn, err = opts.Connect()
	if err != nil {
		return err
	}

	nc, err = nats.NewEncodedConn(conn, "json")
	if err != nil {
		return err
	}

	return nil
}

// Conn get Connection
func Conn() *nats.EncodedConn {
	return nc
}

func Request(ctx context.Context, subject string, timeout uint16, request interface{}, dest interface{}) error {
	nctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(timeout))
	defer cancel()

	if err := nc.RequestWithContext(nctx, subject, request, dest); err != nil {
		return err
	}
	return nil
}
