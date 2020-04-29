package tests

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"sancap/internal/configs"
	"sancap/internal/routers"
	"testing"
)

func TestPingRoute(t *testing.T) {
	var options struct {
		Config      string `short:"c" long:"config"`
		Environment string `short:"e" long:"environment" default:"development"`
	}
	//p := flags.NewParser(&options, flags.Default)
	//if _, err := p.ParseArgs(os.Args[4:]); err != nil {
	//	log.Panicln(err)
	//}
	options.Environment = "development"

	if options.Config == "" {
		options.Config = "/home/korhan/Projects/Go/sancap/internal/configs/config.yaml"
	}

	if err := configs.Init(options.Config, options.Environment); err != nil {
		log.Panicln(err)
	}

	router := routers.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/register", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())

}
