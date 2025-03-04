package devicemodel

import "Device_Scalable_Backend/src/common"

type DeviceGet struct {
	common.SqlModelDefault
	AtmId             string              `gorm:"atm_id" json:"atm_id"`
	Address           string              `gorm:"address" json:"address"`
	Agency            string              `gorm:"agency" json:"agency"`
	DevicePhoneAccess []DevicePhoneAccess `gorm:"foreignKey:DeviceId;references:Id;" json:"phone_access"`
}

type DevicePhoneAccess struct {
	common.SqlModelDefault
	DeviceId int    `gorm:"device_id"`
	Value    string `gorm:"value"`
}

func (DevicePhoneAccess) TableName() string {
	return "device_phone_access"
}
