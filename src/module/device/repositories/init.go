package devicerepositories

import (
	devicemodel "Device_Scalable_Backend/src/module/device/model"

	"gorm.io/gorm"
)

type DeviceRepository struct {
	sql *gorm.DB
}

func NewDeviceRepository(sqlDevice *gorm.DB) *DeviceRepository {
	return &DeviceRepository{sql: sqlDevice.Table(devicemodel.DeviceTableName)}
}
