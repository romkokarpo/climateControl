package repositories

import (
	"github.com/romkokarpo/climateControl/DAL"
	"github.com/romkokarpo/climateControl/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceRepository struct {
	dbContext *DAL.DbContext
}

func NewDeviceRepository() *DeviceRepository {
	return &DeviceRepository{
		dbContext: DAL.NewDbContext(),
	}
}

func (repository *DeviceRepository) GetAllDevices() []*models.Device {
	ctx := repository.dbContext
	var result []*models.Device

	cursor, err := ctx.Devices.Find(*ctx.MongoContext, bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(*ctx.MongoContext, &result); err != nil {
		panic(err)
	}
	return result
}

func (repository *DeviceRepository) GetDeviceById(id string) *models.Device {
	ctx := repository.dbContext
	hexedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	queryResult := ctx.Devices.FindOne(*ctx.MongoContext, bson.M{"_id": hexedId})
	device := models.Device{}
	queryResult.Decode(device)

	return &device
}

func (repository *DeviceRepository) GetDeviceByNameAndModel(name string, model string) *models.Device {
	ctx := repository.dbContext

	queryResult := ctx.Devices.FindOne(
		*ctx.MongoContext, bson.M{
			"name":  name,
			"model": model,
		},
	)

	device := models.Device{}
	queryResult.Decode(device)

	return &device
}

func (repository *DeviceRepository) GetDevicesByManufacturer(manufacturer string) *models.Device {
	ctx := repository.dbContext

	queryResult := ctx.Devices.FindOne(*ctx.MongoContext, bson.M{"manufacturer": manufacturer})
	device := models.Device{}
	queryResult.Decode(device)

	return &device
}

func (repository *DeviceRepository) UpdateDevice(device *models.Device) bool {
	ctx := repository.dbContext

	updateResult, err := ctx.Devices.UpdateOne(
		*ctx.MongoContext,
		bson.M{"_id": device.ID},
		bson.M{
			"$set": bson.M{
				"name":                   device.Name,
				"model":                  device.Model,
				"dateOfPresentation":     device.DateOfPresentation,
				"manufacturer":           device.Manufacturer,
				"wholesalePrice":         device.WholesalePrice,
				"sellPrice":              device.SellPrice,
				"previousWholesalePrice": device.PreviousWholesalePrice,
				"previousSellPrice":      device.PreviousSellPrice,
				"averageErrorRate":       device.AverageErrorRate,
			},
		},
	)

	if err != nil {
		panic(err)
	}

	return updateResult.ModifiedCount > 0
}

func (repository *DeviceRepository) DeleteDevice(id string) bool {
	ctx := repository.dbContext
	hexedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	deleteResult, err := ctx.Devices.DeleteOne(*ctx.MongoContext, bson.M{"_id": hexedId})
	if err != nil {
		panic(err)
	}

	return deleteResult.DeletedCount > 0
}
