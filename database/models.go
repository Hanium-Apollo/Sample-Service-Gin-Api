package database

type Schedule struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Content string `json:"content"`
}

func (Schedule) TableName() string {
	return "schedule"
}

type ScheduleService interface {
	GetSchedules() []Schedule
	CreateSchedule(schedule Schedule) Schedule
	UpdateSchedule(schedule Schedule) Schedule
	DeleteSchedule(schedule Schedule)
}
