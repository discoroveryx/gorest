package repositories

import (
	"app/dbstorage"
	"app/user/models"

	"gorm.io/gorm"
)

type UserCreateRepository struct {
	conn *gorm.DB
}

func NewUserCreateRepository() UserCreateRepository {
	// db := dbstorage.NewDB(new(dbstorage.MyDb))
	// return UserCreateRepository{conn: db.GetConn()}
	db := new(dbstorage.MyDb)
	conn := dbstorage.NewDB1(db)
	return UserCreateRepository{conn: conn}
}

func (u *UserCreateRepository) UserCreate(name string, email string, password string) models.UserModel {

	user := models.UserModel{
		Name:     name,
		Email:    email,
		Password: password,
		// Ctime:    time.Now(),
	}

	u.conn.Create(&user)
	return user
}
