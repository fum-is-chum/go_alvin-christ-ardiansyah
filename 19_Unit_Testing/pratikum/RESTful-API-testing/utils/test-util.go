package utils

import (
	"restful-api-testing/config"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitEchoTestAPI() (*echo.Echo, *gorm.DB) {
	e := echo.New()

	// load config
	var cfg = config.InitConfigTest()

	// open mysql connection
	dbMysql := config.InitMyqlConn(cfg)

	// migrate db
	config.Migrate(dbMysql)

	return e, dbMysql
}