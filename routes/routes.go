package routes

import (
	"scheduleCRUD/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/schedule", handler.GetSchedules)
		api.POST("/schedule", handler.CreateSchedule)
		api.PATCH("/schedule/:id", handler.UpdateSchedule)
		api.DELETE("/schedule/:id", handler.DeleteSchedule)
	}
	return r
}