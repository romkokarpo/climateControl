package DTO

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmbeddedDevice struct {
	ControlSystemID   primitive.ObjectID `bson:"_id,omitempty"`
	DeviceID          primitive.ObjectID `bson:"deviceId,omitempty"`
	SerialNumber      string             `bson:"serialNumber,omitempty"`
	DateOfManufacture time.Time          `bson:"dateOfManufacture,omitempty"`
	CurrentStatus     string             `bson:"currentStatus,omitempty"`
}
