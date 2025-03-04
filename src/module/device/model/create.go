package devicemodel

import "Device_Scalable_Backend/src/common"

type DeviceCreate struct {
	common.SqlModelDefault
	AtmId       string   `gorm:"atm_id" json:"atm_id"`
	Address     string   `gorm:"address" json:"address"`
	Agency      string   `gorm:"agency" json:"agency"`
	PhoneAccess []string `gorm:"-" json:"phone_access"`
}
