package main

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/schedules", GetSchedules)
		api.POST("/schedules", CreateSchedule)
		api.PATCH("/schedules/:id", UpdateSchedule)
		api.DELETE("/schedules/:id", DeleteSchedule)
	}
}