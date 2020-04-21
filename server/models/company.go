package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type EmbeddedHeadQuarters struct {
	Street         string `bson:"street,omitempty"`
	BuildingNumber string `bson:"buildingNumber,omitempty"`
	OpenHour       string `bson:"openHour,omitempty"`
	ClosingHour    string `bson:"closingHour,omitempty"`
}

type EmbeddedOffice struct {
	Street         string `bson:"street,omitempty"`
	BuildingNumber string `bson:"buildingNumber,omitempty"`
	OpenHour       string `bson:"openHour,omitempty"`
	ClosingHour    string `bson:"closingHour,omitempty"`
}

type EmbeddedGreenhouse struct {
	Street         string               `bson:"street,omitempty"`
	BuildingNumber string               `bson:"buldingNumber,omitempty"`
	ControlSystems []primitive.ObjectID `bson:"controlSystems,omitempty"`
}

type Company struct {
	Name         string                `bson:"name,omitempty"`
	HeadQuarters *EmbeddedHeadQuarters `bson:"headQuarters,omitempty"`
	Offices      []*EmbeddedOffice     `bson:"offices,omitempty"`
	Greenhouses  []*EmbeddedGreenhouse `bson:"greenhouses,omitempty"`
}
