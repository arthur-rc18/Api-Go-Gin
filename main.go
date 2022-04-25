package main

import (
	"github.com/arthur-rc18/Api-Go-Gin/database"
	"github.com/arthur-rc18/Api-Go-Gin/routes"
)

func main() {
	database.DbConnection()
	routes.HandleRequests()
}
