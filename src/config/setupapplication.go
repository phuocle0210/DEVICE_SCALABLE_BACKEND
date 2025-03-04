package config

import (
	"Device_Scalable_Backend/src/common"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TSetupApplication struct {
	Engine *gin.Engine
	Ctx    *common.AppContext
}

func SetupApplication(r *gin.Engine) *TSetupApplication {
	dsn := "root:root@tcp(127.0.0.1:3307)/Test?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}

	appContext := common.NewAppContext(db)

	return &TSetupApplication{
		Engine: r,
		Ctx:    appContext,
	}
}
