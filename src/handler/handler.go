package handler

import (
	"github.com/gin-gonic/gin"
)

func GetSchedules(c *gin.Context) {
	var schedules []Schedule
	db.Find(&schedules)
	c.JSON(200, schedules)
}