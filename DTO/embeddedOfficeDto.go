package DTO

import (
	"climateControl/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmbeddedOfficeDto struct {
	CompanyID  primitive.ObjectID       `bson:"_id,omitempty"`
	NewOffices []*models.EmbeddedOffice `bson:"newOffices,omitempty"`
}
