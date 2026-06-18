package rabbitmq

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ShipmentDeliveredEvent struct {
	NoResi     string `json:"no_resi"`
	TrackingID string `json:"tracking_id"`
	Status     string `json:"status"`
}

// DeliveryCreator adalah interface agar rabbitmq tidak perlu import service langsung
type DeliveryCreator interface {
	CreateDeliveryFromShipment(noResi string, trackingID string) error
}

func ConsumeShipmentDelivered(ch *amqp.Channel, deliveryService DeliveryCreator) {
	msgs, err := ch.Consume(
		"shipment.delivered",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Gagal consume shipment.delivered:", err)
	}

	go func() {
		for d := range msgs {
			var event ShipmentDeliveredEvent
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Println("Warning: gagal unmarshal shipment.delivered:", err)
				continue
			}

			if err := deliveryService.CreateDeliveryFromShipment(event.NoResi, event.TrackingID); err != nil {
				log.Println("Warning: gagal create delivery dari shipment.delivered:", err)
			} else {
				log.Printf("Delivery created dari shipment: %s\n", event.NoResi)
			}
		}
	}()
}