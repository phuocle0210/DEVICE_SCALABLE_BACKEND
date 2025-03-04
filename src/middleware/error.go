package middleware

import (
	"Device_Scalable_Backend/src/common"
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				fmt.Println(err)

				if appError, ok := err.(*common.AppError); ok {
					common.RMessage(c).CatchCommonError(appError)
					return
				}

				common.RMessage(c).ErrorJson(500, "something wrong")
				return
			}
		}()

		c.Next()
	}
}
