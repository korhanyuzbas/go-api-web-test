package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"sancap/internal/configs"
	"time"
)

type User struct {
	gorm.Model
	FirstName   string `json:"first_name" form:"first_name"`
	LastName    string `json:"last_name" form:"last_name"`
	Username    string `json:"username" form:"username" binding:"required" validate:"username-check,required"`
	Password    []byte `json:"password" form:"password" json:"password"`
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
	user.Password = user.HashPassword()
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

func (user *User) GetCredentials(password string) (err error) {
	if err = configs.DB.Where("username = ?", user.Username).Find(&user).Error; err != nil {
		return err
	}
	if user.IsCorrectPassword(password) {
		return nil
	}
	return errors.New("wrong user/password")
}

func (user *User) HashPassword() []byte {
	hashedPass, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		log.Panicln(err)
	}
	return hashedPass
}

func (user *User) IsCorrectPassword(password string) bool {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password)) == nil
}

func (user *User) ChangePassword(password string) bool {

	return false
}
