package route

import (
	"Device_Scalable_Backend/src/common"
	devicecontroller "Device_Scalable_Backend/src/module/device/controllers/gin"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, ctx *common.AppContext) {
	v1 := r.Group("/v1")
	{
		deviceRouter := v1.Group("/devices")
		{
			deviceController := devicecontroller.NewDeviceController()

			deviceRouter.GET("/:id", deviceController.Find(ctx))
			deviceRouter.GET("/", deviceController.Paginate(ctx))
			deviceRouter.POST("/", deviceController.Create(ctx))
			deviceRouter.PUT("/:id", deviceController.UpdateByAtmId(ctx))
			deviceRouter.DELETE("/:id", deviceController.Delete(ctx))
		}
	}
}
