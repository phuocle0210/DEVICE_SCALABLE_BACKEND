package devicerepositories

import (
	"Device_Scalable_Backend/src/common"
	devicemodel "Device_Scalable_Backend/src/module/device/model"
	"context"
	"strconv"
)

func (d *DeviceRepository) Paginate(ctx context.Context) *common.TPaginate {
	page, _ := strconv.Atoi(ctx.Value("page").(string))
	limit, _ := strconv.Atoi(ctx.Value("limit").(string))

	var devices []devicemodel.DeviceGet
	v, _ := common.Paginate(d.sql.Preload("DevicePhoneAccess")).Get(&devices, page, limit)

	return v
}
