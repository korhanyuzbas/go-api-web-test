package user

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"regexp"
	"sancap/internal/routers"
	"sancap/tests"
	"strings"
	"testing"
)

func TestUserLogin(t *testing.T) {
	router := tests.SetupTestRouter()
	w := tests.PerformRequest(
		router,
		"POST",
		"/user/login",
		strings.NewReader(
			routers.CreateDataParams(map[string]string{
				"username": "korhan",
				"password": "123456",
			}),
		),
	)

	bodyReader, _ := ioutil.ReadAll(w.Body)
	rgx := regexp.MustCompile("<title>(.*?)</title>")

	fmt.Println(string(bodyReader))
	if rgx.FindString(string(bodyReader)) != "" {
		assert.Equal(t, rgx.FindStringSubmatch(string(bodyReader))[1], "Title")
	}
	assert.Equal(t, w.Code, 200)
}
