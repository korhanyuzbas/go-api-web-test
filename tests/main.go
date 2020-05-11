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
	"sancap/internal/models"
	"sancap/internal/routers"
)

func PerformRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var TestConfig *configs.Option

func createEnvironmentConfig() {
	var options struct {
		Config      string `short:"c" long:"config"`
		Environment string `short:"e" long:"environment" default:"test"`
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
	TestConfig = configs.AppConfig
}

func initDB() {
	var err error
	configs.DB, err = gorm.Open("postgres", configs.DbURL(&configs.DBConfig{
		Host:     TestConfig.Database.Host,
		Port:     TestConfig.Database.Port,
		User:     TestConfig.Database.User,
		Name:     TestConfig.Database.Name,
		Password: TestConfig.Database.Password,
	}))

	configs.DB.AutoMigrate(&models.User{}, &models.UserVerification{})
	if err != nil {
		log.Panicln(err)
	}
}
func Setup() {
	createEnvironmentConfig()
	initDB()
}

func TearDown() {
	configs.DB.DropTableIfExists(&models.User{}, &models.UserVerification{})
	//configs.DB.Close()
}

func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return routers.SetupRouter()
}
