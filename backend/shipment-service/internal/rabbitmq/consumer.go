package rabbitmq

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PickupCompletedEvent struct {
	TrackingNumber string `json:"tracking_number"`
	OrderID        string `json:"order_id"`
	Status         string `json:"status"`
}

// ShipmentCreator adalah interface agar rabbitmq tidak perlu import service langsung
type ShipmentCreator interface {
	CreateShipmentFromPickup(trackingNumber string, orderID string) error
}

func ConsumePickupCompleted(ch *amqp.Channel, shipmentService ShipmentCreator) {
	msgs, err := ch.Consume(
		"pickup.completed",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Gagal consume pickup.completed:", err)
	}

	go func() {
		for d := range msgs {
			var event PickupCompletedEvent
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Println("Warning: gagal unmarshal pickup.completed:", err)
				continue
			}

			if err := shipmentService.CreateShipmentFromPickup(event.TrackingNumber, event.OrderID); err != nil {
				log.Println("Warning: gagal create shipment dari pickup.completed:", err)
			} else {
				log.Printf("Shipment created dari pickup: %s\n", event.TrackingNumber)
			}
		}
	}()
}