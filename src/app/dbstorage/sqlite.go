package dbstorage

import (
	"app/user/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MyDbInter interface {
	NewConn() *gorm.DB
}

func NewDB1(ss MyDbInter) *gorm.DB {
	fmt.Println("\nNewDB1\n")
	return ss.NewConn()
}

type MyDb struct{}

func (d *MyDb) NewConn() *gorm.DB {
	fmt.Println("\n")
	fmt.Println("I am a NewConn")
	fmt.Println("\n")
	conn, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	conn.AutoMigrate(&models.UserModel{})

	return conn
}
