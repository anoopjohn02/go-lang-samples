package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Alerts struct {
	ID          primitive.ObjectID `bson:"_id"`
	DeviceId    string
	Type        string
	Severity    string
	Key         string
	Description string
	DateTime    time.Time
}
