package deviceusecases

import (
	devicemodel "Device_Scalable_Backend/src/module/device/model"
	devicerepositories "Device_Scalable_Backend/src/module/device/repositories"
	"context"
)

type DeviceDeleteRepository interface {
	Delete(ctx context.Context,
		cond map[string]interface{},
		data *devicemodel.DeviceDelete,
	) error
}

type DeviceDeleteUseCase struct {
	repo DeviceDeleteRepository
}

func NewDeviceDeleteUseCase(deviceRepo *devicerepositories.DeviceRepository) *DeviceDeleteUseCase {
	return &DeviceDeleteUseCase{repo: deviceRepo}
}

func (d *DeviceDeleteUseCase) DeleteByAtmId(
	ctx context.Context,
	id string,
	data *devicemodel.DeviceDelete,
) error {
	if err := d.repo.Delete(ctx, map[string]interface{}{"atm_id": id}, data); err != nil {
		return err
	}
	return nil
}
