package rabbitmq

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	Channel *amqp.Channel
}

// PublishShipmentDelivered dipublish saat status shipment jadi DELIVERED
func (p *Publisher) PublishShipmentDelivered(noResi string, trackingID string) error {
	body, _ := json.Marshal(map[string]interface{}{
		"no_resi":    noResi,
		"tracking_id": trackingID,
		"status":     "DELIVERED",
	})

	return p.Channel.Publish(
		"",
		"shipment.delivered",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

// PublishShipmentCreated dipublish saat shipment baru dibuat (untuk tracking/notification)
func (p *Publisher) PublishShipmentCreated(noResi string, trackingID string, originCity string, destinationCity string) error {
	body, _ := json.Marshal(map[string]interface{}{
		"no_resi":          noResi,
		"tracking_id":      trackingID,
		"origin_city":      originCity,
		"destination_city": destinationCity,
		"status":           "IN_TRANSIT",
	})

	return p.Channel.Publish(
		"",
		"shipment.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}