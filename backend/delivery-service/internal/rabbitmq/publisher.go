package rabbitmq

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	Channel *amqp.Channel
}

// PublishDeliveryCompleted dipublish saat paket berhasil diterima penerima
func (p *Publisher) PublishDeliveryCompleted(trackingID string, noResi string) error {
	body, _ := json.Marshal(map[string]interface{}{
		"tracking_id": trackingID,
		"no_resi":     noResi,
		"status":      "DELIVERED",
	})

	return p.Channel.Publish(
		"",
		"delivery.completed",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}