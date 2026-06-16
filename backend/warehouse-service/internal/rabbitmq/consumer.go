package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type PickupCompletedEvent struct {
	TrackingNumber string `json:"tracking_number"`
	OrderID        string `json:"order_
	id"`
	Status         string `json:"status"`
}

// WarehouseLogCreator adalah interface kecil agar paket rabbitmq tidak perlu
// import paket service secara langsung (menghindari import cycle).
type WarehouseLogCreator interface {
	CreateLogFromPickup(trackingNumber string) error
}

func ConsumePickupCompleted(
	ch *amqp091.Channel,
	warehouseService WarehouseLogCreator,
) {

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
		log.Fatal(err)
	}

	go func() {

		for d := range msgs {

			var event PickupCompletedEvent

			err := json.Unmarshal(d.Body, &event)

			if err != nil {
				continue
			}

			err = warehouseService.CreateLogFromPickup(event.TrackingNumber)

			if err != nil {
				log.Println("Warning: gagal membuat warehouse log dari pickup.completed:", err)
			}
		}
	}()
}
