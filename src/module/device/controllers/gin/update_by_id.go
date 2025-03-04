package devicecontroller

import (
	"Device_Scalable_Backend/src/common"
	devicemodel "Device_Scalable_Backend/src/module/device/model"
	devicerepositories "Device_Scalable_Backend/src/module/device/repositories"
	deviceusecases "Device_Scalable_Backend/src/module/device/usecases"

	"github.com/gin-gonic/gin"
)

func (d *DeviceController) UpdateByAtmId(ctx *common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		deviceRepo := devicerepositories.NewDeviceRepository(ctx.GetMainDBConnect())
		deviceUseCase := deviceusecases.NewDeviceUpdateUseCase(deviceRepo)

		var device *devicemodel.DeviceUpdate
		if err := c.ShouldBind(&device); err != nil {
			common.Error(err, "Khong lay duoc du lieu").Throw()
		}

		r := devicemodel.DeviceUpdate{
			Address:     device.Address,
			Agency:      device.Agency,
			PhoneAccess: device.PhoneAccess,
		}

		err := deviceUseCase.UpdateByAtmId(c, c.Param("id"), &r)
		if err != nil {
			common.Error(err, "Sua khong thanh cong").Throw()
		}
		common.RMessage(c).SuccessJson(200, "OK", nil)
	}
}
