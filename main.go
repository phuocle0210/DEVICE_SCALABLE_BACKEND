package main

import (
	"Device_Scalable_Backend/src/config"
	"Device_Scalable_Backend/src/middleware"
	route "Device_Scalable_Backend/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := config.SetupApplication(gin.Default())

	app.Engine.Use(middleware.HandleError())
	route.Routes(app.Engine, app.Ctx)
	app.Engine.Run()
}
