package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceRealTimeData struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	DateTaken       time.Time          `bson:"dateTaken,omitempty"`
	ControlSystemID primitive.ObjectID `bson:"controlSystemId,omitempty"`
	DeviceID        primitive.ObjectID `bson:"deviceId,omitempty"`
	Value           string             `bson:"value,omitempty"`
	Unit            string             `bson:"unit,omitempty"`
}
