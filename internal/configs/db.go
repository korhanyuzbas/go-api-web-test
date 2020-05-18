package configs

import (
	"fmt"
)

//var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Name,
		dbConfig.Password,
	)
}
