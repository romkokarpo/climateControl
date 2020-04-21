package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	Name                   string             `bson:"name,omitempty"`
	Model                  string             `bson:"model,omitempty"`
	DateOfPresentation     time.Time          `bson:"dateOfPresentation,omitempty"`
	Manufacturer           string             `bson:"manufacturer,omitempty"`
	WholesalePrice         float32            `bson:"wholesalePrice,omitempty"`
	SellPrice              float32            `bson:"sellPrice,omitempty"`
	PreviousWholesalePrice float32            `bson:"previousWholesalePrice,omitempty"`
	PreviousSellPrice      float32            `bson:"previousSellPrice,omitempty"`
	AverageErrorRate       float32            `bson:"averageErrorRate,omitempty"`
}
