// package database

// import (
// 	"fmt"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func init() {
// 	godotenv.Load()
// 	var err error
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 	os.Getenv("DB_USER"),
// 	os.Getenv("DB_PASSWORD"),
// 	os.Getenv("DB_HOST"),
// 	os.Getenv("DB_PORT"),
// 	os.Getenv("DB_NAME"))

// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	db.autoMigrate(&Schedule{})
// }

package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type Schedule struct {
	ID uint `json:"id" gorm:"primary_key"`
	Content string `json:"content"`
}

func getEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func getDBClient() {
	var (
		host = getEnvVar("DB_HOST")
		port = getEnvVar("DB_PORT")
		user = getEnvVar("DB_USER")
		password = getEnvVar("DB_PASSWORD")
		dbname = getEnvVar("DB_NAME")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host,
			port,
			user,
			password,
			dbname,
	)
	fmt.Println(conn)

	db, err = gorm.Open(mysql.Open(conn), &gorm.Config{})

}