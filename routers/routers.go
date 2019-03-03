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
	Routers.GET("/b/cards/review/:id", controllers.AllCardReview)
	Routers.GET("/b/cards/overdue/:id", controllers.AllCardChangeDueDate)
}
