package handlers

import (
	"github.com/jinzhu/gorm"
)

type BaseHandler struct {
	DB *gorm.DB
}
