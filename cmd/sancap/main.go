package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"sancap/internal/configs"
	"sancap/internal/middlewares"
	"sancap/internal/routers"
)

var err error

func runServer() (err error) {
	authMiddleware := middlewares.Authentication()
	router := routers.SetupRouter(authMiddleware) // setup routers

	return router.Run()
}

func main() {
	configs.DB, err = gorm.Open("postgres", configs.DbURL(configs.BuildDBConfig()))
	if err != nil {
		log.Panicln(err)
	}

	defer configs.DB.Close()

	log.Panicln(runServer())
}
