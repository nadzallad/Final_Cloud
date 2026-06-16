package rabbitmq

import (
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	Channel *amqp091.Channel
}

func (p *Publisher) Publish(
	queue string,
	payload interface{},
) error {

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return p.Channel.Publish(
		"",
		queue,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}