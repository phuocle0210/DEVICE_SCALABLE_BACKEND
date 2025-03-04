package deviceusecases

import (
	"Device_Scalable_Backend/src/common"
	devicerepositories "Device_Scalable_Backend/src/module/device/repositories"
	"context"
)

type DevicePaginateRepository interface {
	Paginate(ctx context.Context) *common.TPaginate
}

type DevicePaginateUseCase struct {
	repo DevicePaginateRepository
}

func NewDevicePaginateUseCase(deviceRepo *devicerepositories.DeviceRepository) *DevicePaginateUseCase {
	return &DevicePaginateUseCase{repo: deviceRepo}
}

func (d *DevicePaginateUseCase) Paginate(ctx context.Context) *common.TPaginate {
	return d.repo.Paginate(ctx)
}
