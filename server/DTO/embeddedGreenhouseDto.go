package DTO

import (
	"github.com/romkokarpo/climateControl/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmbeddedGreenhouseDto struct {
	CompanyID   primitive.ObjectID           `bson:"_id,omitempty"`
	Greenhouses []*models.EmbeddedGreenhouse `bson:"greenhouses,omitempty"`
}
