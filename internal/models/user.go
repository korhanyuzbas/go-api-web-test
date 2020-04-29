package models

import (
	"github.com/jinzhu/gorm"
	"sancap/internal/configs"
	"time"
)

type User struct {
	gorm.Model
	FirstName   string `json:"first_name" form:"first_name"`
	LastName    string `json:"last_name" form:"last_name"`
	Username    string `json:"username" form:"username" binding:"required" validate:"username-check,required"`
	Password    string `json:"password" form:"password" json:"password"`
	IsActive    bool   `gorm:"default:false"`
	IsSuperUser bool   `gorm:"default:false"`
	DOB         *time.Time
}

type StudentModel struct {
	gorm.Model
	UserID       int
	User         User
	ProfilePhoto string
}

func (user *User) TableName() string {
	return "user"
}

func (user *User) Create() (err error) {
	if err = configs.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) IsUsernameAvailable(username string) bool {
	if configs.DB.Where("username = ?", username).First(user).RecordNotFound() {
		return true
	}
	return false
}

func (user *User) GetByName(username string) (err error) {
	if err = configs.DB.Where("username = ?", username).Find(&user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) GetCredentials(username string, password string) (err error) {
	if err = configs.DB.Where("username = ? AND password = ?", username, password).Find(&user).Error; err != nil {
		return err
	}
	return nil
}
