package config

import (
	"Device_Scalable_Backend/src/common"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TSetupApplication struct {
	Engine *gin.Engine
	Ctx    *common.AppContext
}

func SetupApplication(r *gin.Engine) *TSetupApplication {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("MYSQL_USER_ROOT")
	pass := os.Getenv("MYSQL_PASSWORD_ROOT")
	host := os.Getenv("MYSQL_HOST")
	database := os.Getenv("MYSQL_DATABASE")
	fmt.Println(user, pass, host, database)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", user, pass, host, database)
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
