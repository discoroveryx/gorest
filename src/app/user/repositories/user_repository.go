package repositories

import (
	"app/dbstorage"
	"app/user/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository() UserRepository {
	// db := dbstorage.NewDB(new(dbstorage.MyDb))
	// return UserRepository{conn: db.GetConn()}
	db := new(dbstorage.MyDb)
	conn := dbstorage.NewDB1(db)
	return UserRepository{conn: conn}
}

func (u *UserRepository) UserExistsByName(name string) (models.UserModel, error) {
	user := models.UserModel{}

	result := u.conn.Where("name = ?", name).Take(&user)
	// fmt.Println(result.Error)

	return user, result.Error
}

func (u *UserRepository) GetUserById(id uint) (models.UserModel, error) {
	user := models.UserModel{}

	result := u.conn.Where("id = ?", id).First(&user)
	// fmt.Println(result.Error)

	return user, result.Error
}

func (u *UserRepository) UserExistsByEmail(email string) bool {
	result := u.conn.Where("email = ?", email).First(&models.UserModel{})

	fmt.Println(result.Error)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
