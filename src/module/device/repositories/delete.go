package devicerepositories

import (
	devicemodel "Device_Scalable_Backend/src/module/device/model"
	"context"

	"gorm.io/gorm"
)

func (d *DeviceRepository) Delete(
	ctx context.Context,
	cond map[string]interface{},
	data *devicemodel.DeviceDelete,
) error {
	return d.sql.Transaction(func(tx *gorm.DB) error {
		var d *devicemodel.DeviceDelete
		if err := tx.Table(devicemodel.DeviceTableName).Where(cond).Find(&d).Error; err != nil {
			return err
		}

		if err := tx.Table(devicemodel.DevicePhoneAccessName).
			Where("device_id = ?", d.Id).
			Delete(nil).Error; err != nil {
			return err
		}

		if err := tx.Table(devicemodel.DeviceTableName).Where(cond).Delete(&data).Error; err != nil {
			return err
		}

		return nil
	})
}
