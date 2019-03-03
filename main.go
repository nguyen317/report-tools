package main

import (
	"./config"
	"./connectAPI"
)

func main() {
	Config := config.ReadConfig()

	// gin.SetMode(gin.DebugMode)
	// r := routers.Routers
	// routers.SetupRouters()
	connectAPI.UpdateDataOnDB(Config.App.Keyapp, Config.App.Token, Config.App.Idboard)
	// r.Run(":" + Config.Server.Port)
}
