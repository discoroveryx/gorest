package dbstorage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenPostgres(host string, port string, user string, name string, tz string) *gorm.DB {
	// TODO check test env
	db, err := connectPostgres(host, port, user, name, tz)
	if err != nil {
		createDatabaseIfNotExists(host, port, user, name, tz)
		db, _ := connectPostgres(host, port, user, name, tz)
		return db
	}

	return db
}

func connectPostgres(host string, port string, user string, name string, tz string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		host, user, name, port, tz,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}

func createDatabaseIfNotExists(host string, port string, user string, name string, tz string) {
	dsn := fmt.Sprintf(
		"host=%s user=%s port=%s sslmode=disable TimeZone=%s",
		host, user, port, tz,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// check if db exists
	type Result struct{ Datname string }

	var result Result

	stmt := fmt.Sprintf("SELECT datname FROM pg_database WHERE datname = '%s';", name)
	db.Raw(stmt).Scan(&result)
	// fmt.Print("\nResult", result)

	if result == (Result{}) {
		query := fmt.Sprintf("CREATE DATABASE %s;", name)
		// fmt.Print(query)
		db.Exec(query)
	}
}
