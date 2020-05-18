package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sancap/internal/configs"
	"sancap/internal/handlers"
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
}

func SetupTest() {
	gin.SetMode(gin.TestMode)
	createEnvironmentConfig()
}
func SetupTestRouter() (handlers.BaseHandler, *gin.Engine) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)
	db.AutoMigrate(models.User{})
	return routers.SetupRouter(db)
}

func CreateTestParams(params map[string]string) string {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	return values.Encode()
}
