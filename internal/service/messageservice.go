package service

import "com/anoop/examples/internal/models"

type MessageService interface {
	Publish(message models.IotoMessage)
	Sub(topic string)
}
