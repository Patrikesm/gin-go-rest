package main

import (
	"github.com/patrike-miranda/gin-go-rest/database"
	"github.com/patrike-miranda/gin-go-rest/routes"
)

func main() {
	database.ConnectDB()

	routes.HandleRequests()
}
