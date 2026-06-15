package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type PaymentSuccessEvent struct {
	OrderID string `json:"order_id"`
}

// PickupCreator adalah interface kecil agar paket rabbitmq tidak perlu
// import paket service secara langsung (menghindari import cycle).
type PickupCreator interface {
	CreatePickupFromPayment(orderID string) error
}

func ConsumePaymentSuccess(
	ch *amqp091.Channel,
	pickupService PickupCreator,
) {

	msgs, err := ch.Consume(
		"payment.success",
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

			var event PaymentSuccessEvent

			err := json.Unmarshal(d.Body, &event)

			if err != nil {
				continue
			}

			err = pickupService.CreatePickupFromPayment(event.OrderID)

			if err != nil {
				log.Println("Warning: gagal membuat pickup dari payment.success:", err)
			}
		}
	}()
}
