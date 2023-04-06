package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Alerts struct {
	Id          primitive.ObjectID
	DeviceId    string
	Type        string
	Severity    string
	Key         string
	Description string
	DateTime    time.Time
}
