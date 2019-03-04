package routers

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

var Routers *gin.Engine

func init() {
	Routers = gin.Default()
}

func SetupRouters() {
	Routers.GET("/b/cards/review", controllers.AllCardReview)
	Routers.GET("/b/cards/change-due", controllers.AllCardChangeDueDate)
}
