package main

import (
	"./routers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := routers.Routers
	routers.SetupRouters()
	// go connectAPI.UpdateDataOnDB()
	r.Run(":8080")

}
