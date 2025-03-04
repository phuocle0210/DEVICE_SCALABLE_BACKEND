package devicerepositories

import (
	"Device_Scalable_Backend/src/component"
	devicemodel "Device_Scalable_Backend/src/module/device/model"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func (d *DeviceRepository) Create(ctx context.Context, data devicemodel.DeviceCreate) error {
	uuid := component.UUID()
	data.AtmId = uuid.UUIDText

	d.sql.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&data)

		if result.Error != nil {
			return result.Error
		}

		if len(data.PhoneAccess) > 0 {
			fmt.Println(len(data.PhoneAccess))
			for i := range data.PhoneAccess {
				tx.Table("device_phone_access").Create(&map[string]interface{}{
					"device_id": data.Id,
					"value":     data.PhoneAccess[i],
				})
			}
		}

		return nil
	})

	return nil
}
