package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"sancap/internal/configs"
	"sancap/internal/utils"
	"time"
)

type User struct {
	gorm.Model  `faker:"-"`
	FirstName   string     `json:"first_name" form:"first_name" faker:"first_name"`
	LastName    string     `json:"last_name" form:"last_name" faker:"last_name"`
	Username    string     `json:"username" form:"username" binding:"required" validate:"username-check,required" faker:"username"`
	Password    []byte     `json:"password" form:"password" json:"password" faker:"-"`
	IsActive    bool       `gorm:"default:false" faker:"-"`
	IsSuperUser bool       `gorm:"default:false" faker:"-"`
	DOB         *time.Time `faker:"-"`
}

type UserVerification struct {
	gorm.Model
	UserID     int
	User       User
	Code       string
	Verified   bool
	ExpiryDate *time.Time
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

func (user *User) Create() bool {
	user.Password = user.HashPassword()
	if err := configs.DB.Create(user).Error; err != nil {
		log.Panicln(err.Error())
		return false
	}
	return true
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

func (user *User) GetCredentials(password string) error {
	if err := configs.DB.Where("username = ?", user.Username).Find(user).Error; err != nil {
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
	if err := configs.DB.Model(&user).Update("password", utils.HashPassword(password)).Error; err != nil {
		log.Panicln(err)
		return false
	}
	return true
}
