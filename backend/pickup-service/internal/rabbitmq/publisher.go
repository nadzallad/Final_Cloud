package rabbitmq

import (
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	Channel *amqp091.Channel
}

func (p *Publisher) PublishPickupCompleted(trackingNumber string, orderID string) error {
	body, _ := json.Marshal(
		map[string]interface{}{
			"tracking_number": trackingNumber,
			"order_id":        orderID,
			"status":          "PICKED_UP",
		},
	)

	return p.Channel.Publish(
		"",
		"pickup.completed",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
