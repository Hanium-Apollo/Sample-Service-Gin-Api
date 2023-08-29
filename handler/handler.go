package handler

import (
	"scheduleCRUD/database"

	"github.com/gin-gonic/gin"
)

func GetSchedules(c *gin.Context) {
	var schedules []database.Schedule
	database.DB.Find(&schedules)
	c.JSON(200, gin.H{"data": schedules})
}

func CreateSchedule(c *gin.Context) {
	var schedule database.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&schedule)
	c.JSON(200, gin.H{"data": schedule})
}

func UpdateSchedule(c *gin.Context) {
	id := c.Param("id")
	var schedule database.Schedule
	if database.DB.First(&schedule, id).Error != nil {
		c.JSON(404, gin.H{"error": "Schedule not found"})
		return
	}
	database.DB.Save(&schedule)
	c.JSON(200, gin.H{"data": schedule})
}

func DeleteSchedule(c *gin.Context) {
	id := c.Param("id")
	var schedule database.Schedule
	if database.DB.First(&schedule, id).Error != nil {
		c.JSON(404, gin.H{"error": "Schedule not found"})
		return
	}
	database.DB.Delete(&schedule)
	c.JSON(200, gin.H{"data": true})
}