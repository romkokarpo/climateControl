package DTO

import (
	"github.com/romkokarpo/climateControl/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmbeddedOfficeDto struct {
	CompanyID primitive.ObjectID       `bson:"_id,omitempty"`
	Offices   []*models.EmbeddedOffice `bson:"newOffices,omitempty"`
}
