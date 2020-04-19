package repositories

import "climateControl/DAL"

type DeviceRepository struct {
	dbContext *DAL.DbContext
}

func NewDeviceRepository() *DeviceRepository {
	return &DeviceRepository{
		dbContext: DAL.NewDbContext(),
	}
}

// func (repository *DeviceRepository) GetAllDevices() []*models.Device {

// }
