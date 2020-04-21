package repositories

import (
	"climateControl/DAL"
	"climateControl/DTO"
	"climateControl/models"
	"log"

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

func (repository *ControlSystemRepository) UpdateSystemDevice(deviceDto DTO.EmbeddedDeviceDto) bool {
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

func (repository *ControlSystemRepository) DeleteSystemDevices(devices []*DTO.EmbeddedDeviceDto) int64 {
	ctx := repository.dbContext
	deviceIds := getDeviceIdsSlice(devices)

	deleteResult, err := ctx.ControlSystems.UpdateOne(
		*ctx.MongoContext,
		bson.M{
			"_id": devices[0].ControlSystemID,
			"devices.deviceId": bson.M{
				"$in": deviceIds,
			},
		},
		bson.M{
			"$pull": bson.M{
				"devices": bson.M{
					"deviceId": bson.M{
						"$in": deviceIds,
					},
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}

	return deleteResult.ModifiedCount
}

func getDeviceIdsSlice(devices []*DTO.EmbeddedDeviceDto) []primitive.ObjectID {
	var result []primitive.ObjectID
	result = make([]primitive.ObjectID, len(devices), len(devices))

	for index, value := range devices {
		result[index] = value.DeviceID
		log.Println(value.DeviceID.String())
	}

	return result
}
