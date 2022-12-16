package main

import (
	"app/dbstorage"
	"fmt"

	"transport"
)

func main() {
	// Init Database
	fmt.Println("I am the main")
	db := new(dbstorage.MyDb)
	dbstorage.NewDB1(db)

	// db := dbstorage.NewDB(new(dbstorage.MyDb))
	// db.Migrate()

	// Init http
	router := transport.SetupRouter()
	router.Run(":8080")
}
