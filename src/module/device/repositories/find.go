package devicerepositories

import (
	devicemodel "Device_Scalable_Backend/src/module/device/model"
	"context"
)

func (d *DeviceRepository) Find(ctx context.Context, id string) (*devicemodel.DeviceGet, error) {
	var deviceFind *devicemodel.DeviceGet
	q := d.sql.Preload("DevicePhoneAccess").Where("atm_id = ?", id).Find(&deviceFind)
	if q.Error != nil {
		return nil, q.Error
	}
	return deviceFind, nil
}
