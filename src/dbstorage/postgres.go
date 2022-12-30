package dbstorage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenPostgres(dbhost string, dbuser string, dbname string) *gorm.DB {
	// fmt.Println("OpenPostgres", dbname)
	// TODO check test env
	CreateDatabaseIfNotExists(dbhost, dbuser, dbname)

	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Novosibirsk",
		dbhost, dbuser, dbname,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func CreateDatabaseIfNotExists(dbhost string, dbuser string, dbname string) {
	// fmt.Println("CreateDatabaseIfNotExists", dbname)
	dsn := fmt.Sprintf(
		"host=%s user=%s port=5432 sslmode=disable TimeZone=Asia/Novosibirsk",
		dbhost, dbuser,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// check if db exists
	type Result struct{ Datname string }

	var result Result

	stmt := fmt.Sprintf("SELECT datname FROM pg_database WHERE datname = '%s';", dbname)
	db.Raw(stmt).Scan(&result)
	// fmt.Print("\nResult", result)

	if result == (Result{}) {
		query := fmt.Sprintf("CREATE DATABASE %s;", dbname)
		// fmt.Print(query)
		db.Exec(query)
	}
}
