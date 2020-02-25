package rabbitmq

import (
	"errors"
	"fmt"

	"github.com/streadway/amqp"

	"github.com/julienbreux/triggy/internal/adapter"
	"github.com/julienbreux/triggy/internal/adapter/rabbitmq/parameter"
	"github.com/julienbreux/triggy/internal/provider/action"
	"github.com/julienbreux/triggy/internal/provider/trigger"
)

const (
	name    = "rabbitmq"
	version = "0.0.1"
)

type instance struct {
	configured bool
	param      *parameter.Parameter

	conn    *amqp.Connection
	channel *amqp.Channel
}

// New creates a new RabbitMQ adapter
func New() adapter.Adapter {
	return &instance{}
}

// Name returns the adapter name
func (a *instance) Name() string {
	return name
}

// Version returns the adapter version
func (a *instance) Version() string {
	return version
}

// Configure configures the adapter
func (a *instance) Configure(params map[string]interface{}) error {
	p, err := parameter.New(params)
	if err != nil {
		return err
	}
	a.param = p

	return nil
}

// Connect connects the adapter to server
func (a *instance) Connect() error {
	if !a.configured {
		return errors.New("adapter: adapter not configured")
	}

	var err error
	a.conn, err = amqp.Dial(a.param.URL().String())
	if err != nil {
		return fmt.Errorf("adapter: %s: %s", name, err)
	}

	// go func() {
	// 	fmt.Printf("closing: %s", <-a.conn.NotifyClose(make(chan *amqp.Error)))
	// }()

	a.channel, err = a.conn.Channel()
	if err != nil {
		return fmt.Errorf("adapter: %s: channel: %s", name, err)
	}

	return nil
}

// Disconnect disconnects the adapter from server
func (a *instance) Disconnect() error {
	if a.conn != nil {
		return a.conn.Close()
	}

	return nil
}

// HealthCheck checks the health of the adapter
func (a *instance) HealthCheck() error {
	return nil
}

// Actions
func (a *instance) Actions() map[string]action.Action {
	return nil
}

// Trigger
func (a *instance) Trigger() map[string]trigger.Trigger {
	return nil
}
