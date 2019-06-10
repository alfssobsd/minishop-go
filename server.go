package main

import (
	mongo_config "github.com/alfssobsd/minishop/config"
	_goodsController "github.com/alfssobsd/minishop/entrypoints/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))

	mgoSession := mongo_config.MakeMongoConnection()
	_goodsController.GoodsRoutes(e, mgoSession)
	e.Logger.Fatal(e.Start(":1323"))
}
