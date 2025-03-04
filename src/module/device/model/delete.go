package devicemodel

import "Device_Scalable_Backend/src/common"

type DeviceDelete struct {
	common.SqlModelDefault
	AtmId string `gorm:"atm_id" json:"atm_id"`
}
