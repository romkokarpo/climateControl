package repositories

import (
	"climateControl/DAL"
	"climateControl/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ControlSystemRepository struct {
	dbContext *DAL.DbContext
}

func NewControlSystemRepository() *ControlSystemRepository {
	return &ControlSystemRepository{
		dbContext: DAL.NewDbContext(),
	}
}

func (repository *ControlSystemRepository) GetAllControlSystems() []*models.ControlSystem {
	ctx := repository.dbContext
	var controlSystems []*models.ControlSystem

	cursor, err := ctx.ControlSystems.Find(
		*ctx.MongoContext,
		bson.M{},
	)
	if err != nil {
		panic(err)
	}

	if err := cursor.All(*ctx.MongoContext, &controlSystems); err != nil {
		panic(err)
	}

	return controlSystems
}

func (repository *ControlSystemRepository) GetControlSystemById(id string) *models.ControlSystem {
	ctx := repository.dbContext
	systemId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	queryResult := ctx.ControlSystems.FindOne(*ctx.MongoContext, bson.M{"_id": systemId})
	controlSystem := models.ControlSystem{}
	queryResult.Decode(controlSystem)

	return &controlSystem
}
