package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"sancap/internal/configs"
	"sancap/internal/routers"
)

func PerformRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func initConfig() {
	var options struct {
		Config      string `short:"c" long:"config"`
		Environment string `short:"e" long:"environment" default:"development"`
	}
	p := flags.NewParser(&options, flags.Default)
	if _, err := p.ParseArgs(os.Args[4:]); err != nil {
		log.Panicln(err)
	}
	if options.Config == "" {
		options.Config = "/home/korhan/Projects/Go/sancap/internal/configs/config.yaml"
	}

	if err := configs.Init(options.Config, options.Environment); err != nil {
		log.Panicln(err)
	}
	var err error
	config := configs.AppConfig
	configs.DB, err = gorm.Open("postgres", configs.DbURL(&configs.DBConfig{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		User:     config.Database.User,
		Name:     config.Database.Name,
		Password: config.Database.Password,
	}))

	if err != nil {
		log.Panicln(err)
	}

}

func SetupTestRouter() *gin.Engine {
	initConfig()
	return routers.SetupRouter(true)
}
