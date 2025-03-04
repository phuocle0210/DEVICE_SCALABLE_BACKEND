package deviceusecases

import (
	devicemodel "Device_Scalable_Backend/src/module/device/model"
	devicerepositories "Device_Scalable_Backend/src/module/device/repositories"
	"context"
)

type DeviceFindRepository interface {
	Find(ctx context.Context, id string) (*devicemodel.DeviceGet, error)
}

type DeviceFindUseCase struct {
	repo DeviceFindRepository
}

func NewDeviceFindUseCase(deviceRepo *devicerepositories.DeviceRepository) *DeviceFindUseCase {
	return &DeviceFindUseCase{repo: deviceRepo}
}

func (d *DeviceFindUseCase) Find(ctx context.Context, id string) (*devicemodel.DeviceGet, error) {
	r, err := d.repo.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return r, nil
}
