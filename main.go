package main

import (
	"scheduleCRUD/database"
	"scheduleCRUD/routes"
)

func init() {
	database.GetDBClient()
}

func main() {
	r := routes.SetupRoutes()
	r.Run(":8080")
}