package main

import (
	"./connectAPI"
	"./routers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
)

var Config = struct {
	App struct {
		Keyapp   string
		Token    string
		Idboard  string
		Username string
	}
	Server struct {
		Host string
		Port string
	}
	Datasabe struct {
		Name     string
		User     string
		Password string
	}

	Contacts struct {
		Email string
	}
}{}

func main() {
	configor.Load(&Config, "config.yml")
	gin.SetMode(gin.DebugMode)
	r := routers.Routers
	routers.SetupRouters()
	go connectAPI.UpdateDataOnDB(Config.App.Keyapp, Config.App.Token, Config.App.Idboard)
	r.Run(":" + Config.Server.Port)
}
