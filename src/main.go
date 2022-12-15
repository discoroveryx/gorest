package main

import (
	"app/dbstorage"
	"fmt"

	"transport"
)

func main() {
	// Init Database
	fmt.Println("I am the main")
	db := dbstorage.NewDB()
	db.Migrate()

	// Init http
	router := transport.SetupRouter()
	router.Run(":8080")
}
