package deviceusecases

import (
	devicemodel "Device_Scalable_Backend/src/module/device/model"
	devicerepositories "Device_Scalable_Backend/src/module/device/repositories"
	"context"
)

type DeviceCreateRepository interface {
	Create(ctx context.Context, data devicemodel.DeviceCreate) error
}

type DeviceUseCase struct {
	repo DeviceCreateRepository
}

func NewDeviceCreateUseCase(deviceRepo *devicerepositories.DeviceRepository) *DeviceUseCase {
	return &DeviceUseCase{repo: deviceRepo}
}

func (d *DeviceUseCase) Create(ctx context.Context, data devicemodel.DeviceCreate) error {
	if err := d.repo.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
