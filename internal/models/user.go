package models

import (
	"github.com/jinzhu/gorm"
	"sancap/internal/configs"
	"time"
)

type UserModel struct {
	gorm.Model
	FirstName   string
	LastName    string
	Username    string
	Password    string
	IsActive    bool `gorm:"default:false"`
	IsSuperUser bool `gorm:"default:false"`
	DOB         *time.Time
}

type StudentModel struct {
	gorm.Model
	UserID       int
	User         UserModel
	ProfilePhoto string
}

func GetUserCred(user *UserModel, username string, password string) (err error) {
	if err = configs.DB.Where("username = ? AND password = ?", username, password).Find(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByName(user *UserModel, username string) (err error) {
	if err = configs.DB.Where("username = ?", username).Find(&user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *UserModel) (err error) {
	if err = configs.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
