package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"sancap/internal/configs"
	"sancap/internal/models"
	"sancap/internal/routers"
)

func runServer(port string, db *gorm.DB) error {
	_, router := routers.SetupRouter(db) // setup routers
	return router.Run(":" + port)
}

func main() {
	var options struct {
		Config      string `short:"c" long:"config"`
		Environment string `short:"e" long:"environment" default:"development"`
	}
	p := flags.NewParser(&options, flags.Default)
	if _, err := p.Parse(); err != nil {
		log.Panicln(err)
	}

	if options.Config == "" {
		options.Config = "/home/korhan/Projects/Go/sancap/internal/configs/config.yaml"
	}
	if err := configs.Init(options.Config, options.Environment); err != nil {
		log.Panicln(err)
	}

	config := configs.AppConfig

	db, err := gorm.Open("postgres", configs.DbURL(&configs.DBConfig{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		User:     config.Database.User,
		Name:     config.Database.Name,
		Password: config.Database.Password,
	}))

	if err != nil {
		log.Panicln(err)
	}

	defer db.Close()
	db.AutoMigrate(&models.User{}, &models.UserVerification{})
	log.Panicln(runServer(config.HTTP.Port, db))
}
