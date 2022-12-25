package main

import (
	"dbstorage"

	"transport"
)

func main() {
	// Init Database
	db := new(dbstorage.DB)
	db.Connect()
	db.Migrate()

	// Init http
	router := transport.SetupRouter(true)
	router.Run(":8080")
}
