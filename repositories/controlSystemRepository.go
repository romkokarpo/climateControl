package repositories

import (
	"climateControl/DAL"
	"climateControl/DTO"
	"climateControl/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (repository *ControlSystemRepository) CreateNewSystem(name string, devices []*models.EmbeddedDevice) *mongo.InsertOneResult {
	ctx := repository.dbContext
	controlSystem := models.ControlSystem{
		Name:    name,
		Devices: devices,
	}

	insertResult, err := ctx.ControlSystems.InsertOne(*ctx.MongoContext, controlSystem)
	if err != nil {
		panic(err)
	}

	return insertResult
}

func (repository *ControlSystemRepository) UpdateSystemName(id string, newName string) bool {
	ctx := repository.dbContext
	systemId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	updateResult, err := ctx.ControlSystems.UpdateOne(
		*ctx.MongoContext,
		bson.M{"_id": systemId},
		bson.M{"name": newName},
	)
	if err != nil {
		panic(err)
	}

	return updateResult.ModifiedCount > 0
}

func (repository *ControlSystemRepository) UpdateSystemDevice(deviceDto DTO.EmbeddedDevice) bool {
	ctx := repository.dbContext

	updateResult, err := ctx.ControlSystems.UpdateOne(
		*ctx.MongoContext,
		bson.M{"_id": deviceDto.ControlSystemID, "devices.deviceId": deviceDto.DeviceID},
		bson.M{
			"$set": bson.M{
				"devices.$.serialNumber":      deviceDto.SerialNumber,
				"devices.$.dateOfManufacture": deviceDto.DateOfManufacture,
				"devices.$.currentStatus":     deviceDto.CurrentStatus,
				"devices.$.deviceId":          deviceDto.DeviceID,
			},
		},
	)
	if err != nil {
		panic(err)
	}

	return updateResult.ModifiedCount > 0
}

func (repository *ControlSystemRepository) DeleteSystem(id string) bool {
	ctx := repository.dbContext
	sysId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	deleteResult, err := ctx.ControlSystems.DeleteOne(*ctx.MongoContext, bson.M{"_id": sysId})
	if err != nil {
		panic(err)
	}

	return deleteResult.DeletedCount > 0
}

func (repository *ControlSystemRepository) DeleteSystemDevices(devices []*DTO.EmbeddedDevice) int64 {
	ctx := repository.dbContext

	deleteResult, err := ctx.Users.DeleteMany(
		*ctx.MongoContext,
		bson.M{
			"_id": devices[0].ControlSystemID,
			"devices.deviceId": bson.M{
				"$in": getDeviceIdsSlice(devices),
			},
		},
	)
	if err != nil {
		panic(err)
	}

	return deleteResult.DeletedCount
}

func getDeviceIdsSlice(devices []*DTO.EmbeddedDevice) []primitive.ObjectID {
	var result []primitive.ObjectID
	result = make([]primitive.ObjectID, len(devices), len(devices))

	for _, value := range devices {
		result = append(result, value.DeviceID)
	}

	return result
}