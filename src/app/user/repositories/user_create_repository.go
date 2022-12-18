package repositories

import (
	"app/user/models"
	"dbstorage"
	"fmt"
	"myconfig"

	"gorm.io/gorm"
)

type UserCreateRepository struct {
	db *gorm.DB
}

func NewUserCreateRepository() UserCreateRepository {
	// db := dbstorage.NewDB(new(dbstorage.MyDb))
	// return UserCreateRepository{conn: db.GetConn()}
	myconf := myconfig.GetMyConfig()
	dbname := myconf.DBName
	fmt.Println("\nUserCreate dbname\n", dbname)

	db := new(dbstorage.DB)
	cursor := db.Connect()
	return UserCreateRepository{db: cursor}
}

func (u *UserCreateRepository) UserCreate(name string, email string, password string) models.UserModel {
	user := models.UserModel{
		Name:     name,
		Email:    email,
		Password: password,
		// Ctime:    time.Now(),
	}

	u.db.Create(&user)
	return user
}
