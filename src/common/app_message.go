package common

import (
	"github.com/gin-gonic/gin"
)

type AppMessage struct {
	ginCtx  *gin.Context
	Message string
	Data    interface{}
	Status  string
}

func RMessage(c *gin.Context) *AppMessage {
	return &AppMessage{ginCtx: c}
}

func (app *AppMessage) CustonMessage(
	status int,
	message string,
	data interface{},
) {
	app.ginCtx.JSON(status, map[string]interface{}{
		"message": message,
		"data":    data,
	})
}

func (app *AppMessage) SuccessJson(
	status int,
	message string,
	data interface{},
) {
	app.CustonMessage(status, message, data)
}

func (app *AppMessage) ErrorJson(
	status int,
	message string,
	data ...interface{},
) {
	app.ginCtx.AbortWithStatusJSON(status, map[string]interface{}{
		"message": message,
		"data":    data,
	})
}

func (app *AppMessage) CatchCommonError(err *AppError) {
	app.ErrorJson(
		err.ErrorCode,
		err.Message,
	)
}
