package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Measurements struct {
	ID       primitive.ObjectID `bson:"_id"`
	DeviceId string
	Type     string
	Value    int
	Unit     string
	Start    time.Time
	End      time.Time
}
