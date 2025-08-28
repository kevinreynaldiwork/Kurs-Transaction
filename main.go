package main

import (
	"FinalProject/database"
	"FinalProject/routers"
	"fmt"
)

func main() {
	db := database.Connect()
	database.DBMigrate(db)

	// Setup routes
	r := routers.SetupRouter(db)
	for _, ri := range r.Routes() {
		fmt.Println(ri.Method, ri.Path)
	}
	// Run server
	r.Run(":8080")

}
