package dbstorage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenPostgres(dbname string) *gorm.DB {
	// fmt.Println("OpenPostgres", dbname)
	// TODO check test env
	CreateDatabaseIfNotExists(dbname)

	dsn := fmt.Sprintf("host=pg user=main_1 dbname=%s port=5432 sslmode=disable TimeZone=Asia/Novosibirsk", dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func CreateDatabaseIfNotExists(dbname string) {
	// fmt.Println("CreateDatabaseIfNotExists", dbname)
	dsn := "host=pg user=main_1 port=5432 sslmode=disable TimeZone=Asia/Novosibirsk"
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
