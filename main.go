package main

import (
	"./config"
	"./connectAPI"
	"./routers"
	"github.com/gin-gonic/gin"
)

func main() {
	Config := config.ReadConfig()

	gin.SetMode(gin.DebugMode)
	r := routers.Routers
	routers.SetupRouters()
	go connectAPI.UpdateDataOnDB(Config.App.Keyapp, Config.App.Token, Config.App.Idboard)
	r.Run(":" + Config.Server.Port)
}
