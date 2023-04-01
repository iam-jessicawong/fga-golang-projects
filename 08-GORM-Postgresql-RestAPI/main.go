package main

import (
	"gorm_postgresql_restapi/database"
	"gorm_postgresql_restapi/routers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)
}
