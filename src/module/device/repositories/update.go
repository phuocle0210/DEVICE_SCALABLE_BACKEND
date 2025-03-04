package devicerepositories

import (
	devicemodel "Device_Scalable_Backend/src/module/device/model"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func (d *DeviceRepository) Update(
	c context.Context,
	cond map[string]interface{},
	data *devicemodel.DeviceUpdate,
) error {
	return d.sql.Transaction(func(tx *gorm.DB) error {
		var d *devicemodel.DeviceUpdate = new(devicemodel.DeviceUpdate)
		tx.Table(devicemodel.DeviceTableName).Where(cond).Find(&d)

		if err := tx.Table(devicemodel.DeviceTableName).Where(cond).Updates(&data).Error; err != nil {
			return err
		}

		fmt.Println(data.PhoneAccess)

		if len(data.PhoneAccess) > 0 {
			if err := tx.Table(devicemodel.DevicePhoneAccessName).
				Where("device_id = ?", d.Id).
				Delete(nil).Error; err != nil {
				return err
			}

			for i := range data.PhoneAccess {
				tx.Table(devicemodel.DevicePhoneAccessName).
					Create(&map[string]interface{}{
						"device_id": d.Id,
						"value":     data.PhoneAccess[i],
					})
			}
		}

		return nil
	})
}
