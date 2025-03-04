package devicecontroller

import (
	"Device_Scalable_Backend/src/common"
	devicemodel "Device_Scalable_Backend/src/module/device/model"
	devicerepositories "Device_Scalable_Backend/src/module/device/repositories"
	deviceusecases "Device_Scalable_Backend/src/module/device/usecases"

	"github.com/gin-gonic/gin"
)

func (d *DeviceController) Create(ctx *common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var device *devicemodel.DeviceCreate

		if err := c.ShouldBind(&device); err != nil {
			common.Error(err, "Khong lay duoc du lieu").Throw()
		}

		deviceRepo := devicerepositories.NewDeviceRepository(ctx.GetMainDBConnect())
		deviceUseCase := deviceusecases.NewDeviceCreateUseCase(deviceRepo)
		if err := deviceUseCase.Create(c, *device); err != nil {
			common.Error(err, "Tao khong thanh cong").Throw()
		}
		common.RMessage(c).SuccessJson(200, "OK", nil)
	}
}
