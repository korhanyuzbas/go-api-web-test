package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"sancap/internal/configs"
	"sancap/internal/routers"
)

var err error

func main() {
	configs.DB, err = gorm.Open("postgres", configs.DbURL(configs.BuildDBConfig()))
	if err != nil {
		log.Panicln(err)
	}

	defer configs.DB.Close()
	router := routers.SetupRouter()
	router.Run()
}
