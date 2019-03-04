package controllers

import (
	"fmt"
	"strconv"
	"time"

	"strings"

	"../database"
	"../modules"
	"github.com/gin-gonic/gin"
)

func AllCardReview(c *gin.Context) {
	q := c.Request.URL.Query()
	result := strings.Split(q["list"][0], ",")

	var data []interface{}

	fmt.Println(result)
	cards, err := database.GetAllCard()
	if err != nil {

	} else {
		for _, v := range result {
			data = append(data, modules.Filter(cards, func(item modules.MyCard) bool {
				return strings.ToLower(v) == strings.ToLower(item.ListName)
			}))
		}
		c.JSON(200, data)
	}
}

func AllCardChangeDueDate(c *gin.Context) {
	var count int
	q := c.Request.URL.Query()
	if q["time"] != nil {
		count, _ = strconv.Atoi(q["time"][0])

	} else {
		count = 7
	}
	cards, err := database.GetAllCard()
	if err != nil {

	} else {
		now := time.Now()
		c.JSON(200, modules.Filter(cards, func(item modules.MyCard) bool {
			return func() int {
				return int(item.DateLastActivity.Sub(now).Hours() / 24)
			}()+count > 0
		}))
	}
}
