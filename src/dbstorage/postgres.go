package dbstorage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenPostgres(dbname string) *gorm.DB {
	dsn := "host=localhost user=main_1 dbname=" + dbname + " port=5432 sslmode=disable TimeZone=Asia/Novosibirsk"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
