package repositories

import (
	"app/user/models"
	"dbstorage"
	"errors"
	"fmt"
	"myconfig"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	// db := dbstorage.NewDB(new(dbstorage.MyDb))
	// return UserRepository{conn: db.GetConn()}
	myconf := myconfig.GetMyConfig()
	dbname := myconf.DBName
	fmt.Println("\nUserRepository dbname\n", dbname)

	db := new(dbstorage.DB)
	cursor := db.Connect()
	return UserRepository{db: cursor}
}

func (u *UserRepository) UserExistsByName(name string) (models.UserModel, error) {
	user := models.UserModel{}

	result := u.db.Where("name = ?", name).Take(&user)
	// fmt.Println(result.Error)

	return user, result.Error
}

func (u *UserRepository) GetUserById(id uint) (models.UserModel, error) {
	user := models.UserModel{}

	result := u.db.Where("id = ?", id).First(&user)
	// fmt.Println(result.Error)

	return user, result.Error
}

func (u *UserRepository) UserExistsByEmail(email string) bool {
	result := u.db.Where("email = ?", email).First(&models.UserModel{})

	fmt.Println(result.Error)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
