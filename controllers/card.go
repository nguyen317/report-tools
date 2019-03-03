package controllers

import (
	"github.com/gin-gonic/gin"
)

func AllCardReview(c *gin.Context) {
	c.JSON(200, gin.H{"meg": "this api for all card review liat and review done on a week"})
}

func AllCardChangeDueDate(c *gin.Context) {
	c.JSON(200, gin.H{"meg": "this api for all card change due date"})
}
