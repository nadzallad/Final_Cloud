package rabbitmq

import (
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	Channel *amqp091.Channel
}

func (p *Publisher) PublishWarehouseCompleted(trackingNumber string) error {
	body, _ := json.Marshal(
		map[string]interface{}{
			"tracking_number": trackingNumber,
			"status":          "OUT_FOR_SHIPMENT",
		},
	)

	return p.Channel.Publish(
		"",
		"warehouse.completed",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
