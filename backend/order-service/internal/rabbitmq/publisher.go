package rabbitmq

import (
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	Channel *amqp091.Channel
}

func (p *Publisher) PublishPaymentSuccess(
	orderID int,
) error {

	body, _ := json.Marshal(
		map[string]interface{}{
			"order_id": orderID,
		},
	)

	return p.Channel.Publish(
		"",
		"payment.success",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body: body,
		},
	)
}