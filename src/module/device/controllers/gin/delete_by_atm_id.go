package devicecontroller

import (
	"Device_Scalable_Backend/src/common"
	devicemodel "Device_Scalable_Backend/src/module/device/model"
	devicerepositories "Device_Scalable_Backend/src/module/device/repositories"
	deviceusecases "Device_Scalable_Backend/src/module/device/usecases"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (d *DeviceController) Delete(ctx *common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var device *devicemodel.DeviceDelete

		if err := c.ShouldBind(&device); err != nil {
			common.Error(err, "Khong lay duoc du lieu").Throw()
		}

		deviceRepo := devicerepositories.NewDeviceRepository(ctx.GetMainDBConnect())
		deviceUseCase := deviceusecases.NewDeviceDeleteUseCase(deviceRepo)
		fmt.Println(c.Param("id"))
		err := deviceUseCase.DeleteByAtmId(c, c.Param("id"), device)
		if err != nil {
			common.Error(err, "Xoa khong thanh cong").Throw()
		}
		common.RMessage(c).SuccessJson(200, "OK", nil)
	}
}
