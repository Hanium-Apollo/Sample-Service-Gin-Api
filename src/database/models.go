package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Schedule struct {
	ID uint "json:id"
	Content string "json:content"
}

type ScheduleService interface {
	GetSchedules() []Schedule
	CreateSchedule(schedule Schedule) Schedule
	UpdateSchedule(schedule Schedule) Schedule
	DeleteSchedule(schedule Schedule)
}

func getEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}