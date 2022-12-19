package dbstorage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenPostgres(dbname string) *gorm.DB {
	CreateDatabaseIfNotExists(dbname)

	dsn := "host=pg user=main_1 dbname=" + dbname + " port=5432 sslmode=disable TimeZone=Asia/Novosibirsk"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func CreateDatabaseIfNotExists(dbname string) {
	fmt.Print("\nCreateDatabaseIfNotExists\n")
	dsn := "host=pg user=main_1 port=5432 sslmode=disable TimeZone=Asia/Novosibirsk"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// check if db exists
	type Result struct{ Datname string }

	var result Result

	stmt := fmt.Sprintf("SELECT datname FROM pg_database WHERE datname = '%s12222';", dbname)
	db.Raw(stmt).Scan(&result)
	// fmt.Print("\nRes", result)

	if result != (Result{}) {
		query := fmt.Sprintf("CREATE DATABASE %s", dbname)
		// fmt.Print(query)
		db.Raw(query)
	}
}
