package devicecontroller

import (
	"Device_Scalable_Backend/src/common"
	devicerepositories "Device_Scalable_Backend/src/module/device/repositories"
	deviceusecases "Device_Scalable_Backend/src/module/device/usecases"

	"github.com/gin-gonic/gin"
)

func (d *DeviceController) Paginate(ctx *common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("page", c.Query("page"))
		c.Set("limit", c.Query("limit"))

		deviceRepo := devicerepositories.NewDeviceRepository(ctx.GetMainDBConnect())
		deviceUseCase := deviceusecases.NewDevicePaginateUseCase(deviceRepo)
		v := deviceUseCase.Paginate(c)

		common.RMessage(c).SuccessJson(200, "OK", v)
	}
}
