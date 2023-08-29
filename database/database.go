package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func GetEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetDBClient() {
	var (
		host = GetEnvVar("DB_HOST")
		port = GetEnvVar("DB_PORT")
		user = GetEnvVar("DB_USER")
		password = GetEnvVar("DB_PASSWORD")
		dbname = GetEnvVar("DB_NAME")
	)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", 
			user, 
			password, 
			host, 
			port, 
			dbname,
	)
	fmt.Println(dns)

	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(DB)
	fmt.Println(err)
	DB.AutoMigrate(&Schedule{})
}