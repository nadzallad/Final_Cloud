package rabbitmq

import (
	"encoding/json"
	"log"

	"order-service/internal/service"

	"github.com/rabbitmq/amqp091-go"
)

type PaymentSuccessEvent struct {
	OrderID int `json:"order_id"`
}

func ConsumePaymentSuccess(
	ch *amqp091.Channel,
	service *service.OrderService,
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

			json.Unmarshal(
				d.Body,
				&event,
			)

			service.MarkAsPaid(
				event.OrderID,
			)
		}

	}()
}