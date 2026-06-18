package rabbitmq

import (
	"encoding/json"
	"log"

	"notification-service/internal/entity"
	"notification-service/internal/service"

	amqp "github.com/rabbitmq/amqp091-go"
)

type DeliveryCompletedEvent struct {
	NoResi          string `json:"no_resi"`
	TrackingID      string `json:"tracking_id"`
	Event           string `json:"event"`
	Message         string `json:"message"`
	DeliveryAddress string `json:"delivery_address"`
	Status          string `json:"status"`
}

func ConsumeDeliveryCompleted(ch *amqp.Channel, svc *service.NotificationService) {
	_, err := ch.QueueDeclare("delivery.completed", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Gagal declare queue delivery.completed:", err)
	}

	msgs, err := ch.Consume(
		"delivery.completed",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Gagal consume delivery.completed:", err)
	}

	go func() {
		for d := range msgs {
			var event DeliveryCompletedEvent
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Println("Warning: gagal unmarshal delivery.completed:", err)
				continue
			}

			// Tentukan message default kalau kosong
			msg := event.Message
			if msg == "" {
				if event.Status == "DELIVERED" {
					msg = "Paket dengan resi " + event.NoResi + " telah berhasil diterima."
				} else {
					msg = "Pengiriman paket dengan resi " + event.NoResi + " gagal."
				}
			}

			// Map ke entity Notification sesuai struct yang ada
			eventType := event.Event
			if eventType == "" {
				if event.Status == "DELIVERED" {
					eventType = "DELIVERY_COMPLETED"
				} else {
					eventType = "DELIVERY_FAILED"
				}
			}

			notif := entity.Notification{
				NoResi:  event.NoResi,
				Source:  "delivery-service",
				Event:   eventType,
				Message: msg,
			}

			if err := svc.CreateNotification(notif); err != nil {
				log.Println("Warning: gagal simpan notifikasi:", err)
			} else {
				log.Printf("Notifikasi tersimpan: [%s] %s\n", eventType, event.NoResi)
			}
		}
	}()
}
