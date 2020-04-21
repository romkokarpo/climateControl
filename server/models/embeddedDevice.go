package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmbeddedDevice struct {
	DeviceID          primitive.ObjectID `bson:"deviceId,omitempty"`
	SerialNumber      string             `bson:"serialNumber,omitempty"`
	DateOfManufacture time.Time          `bson:"dateOfManufacture,omitempty"`
	CurrentStatus     string             `bson:"currentStatus,omitempty"`
}
