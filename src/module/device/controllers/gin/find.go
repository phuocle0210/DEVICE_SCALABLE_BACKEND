package devicecontroller

import (
	"Device_Scalable_Backend/src/common"
	devicerepositories "Device_Scalable_Backend/src/module/device/repositories"
	deviceusecases "Device_Scalable_Backend/src/module/device/usecases"

	"github.com/gin-gonic/gin"
)

func (d *DeviceController) Find(ctx *common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		deviceRepo := devicerepositories.NewDeviceRepository(ctx.GetMainDBConnect())
		deviceUseCase := deviceusecases.NewDeviceFindUseCase(deviceRepo)
		r, err := deviceUseCase.Find(c, id)
		if err != nil {
			common.Error(err, "Khong lay duoc du lieu").Throw()
		}

		common.RMessage(c).SuccessJson(200, "OK", r)
	}
}
