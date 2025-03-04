package deviceusecases

import (
	devicemodel "Device_Scalable_Backend/src/module/device/model"
	devicerepositories "Device_Scalable_Backend/src/module/device/repositories"
	"context"
)

type DeviceUpdateRepository interface {
	Update(c context.Context,
		cond map[string]interface{},
		data *devicemodel.DeviceUpdate,
	) error
}

type DeviceUpdateUseCase struct {
	repo DeviceUpdateRepository
}

func NewDeviceUpdateUseCase(deviceRepo *devicerepositories.DeviceRepository) *DeviceUpdateUseCase {
	return &DeviceUpdateUseCase{repo: deviceRepo}
}

func (d *DeviceUpdateUseCase) UpdateByAtmId(
	ctx context.Context,
	id string,
	data *devicemodel.DeviceUpdate,
) error {
	if err := d.repo.Update(ctx, map[string]interface{}{"atm_id": id}, data); err != nil {
		return err
	}
	return nil
}
