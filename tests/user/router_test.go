package user

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"regexp"
	"sancap/internal/models"
	"sancap/internal/routers"
	"sancap/tests"
	"strings"
	"testing"
)

var router *gin.Engine

func TestStart(t *testing.T) {
	tests.Setup()
	router = tests.SetupTestRouter()
	t.Run("Login", UserLogin)
	t.Run("Login_MissingParams", UserLoginMissingParams)
	t.Run("UserLogin_WrongCredentials", UserLoginWrongCredentials)
	t.Run("UserRegisterSuccess", UserRegisterSuccess)
	t.Run("UserRegisterMissingFields", UserRegisterMissingFields)
	t.Run("UserRegisterUsernameExists", UserRegisterUsernameExists)
	t.Run("UserMe", UserMe)
	tests.TearDown()
}

func UserLogin(t *testing.T) {

	w := tests.PerformRequest(
		router,
		"GET",
		"/user/login",
		nil,
	)
	assert.Equal(t, w.Code, 200)

	user := models.User{Password: []byte("123456")}
	if err := faker.FakeData(&user); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	if err := user.Create(); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	w = tests.PerformRequest(
		router,
		"POST",
		"/user/login",
		strings.NewReader(
			routers.CreateDataParams(map[string]string{
				"username": user.Username,
				"password": "123456",
			}),
		),
	)

	bodyReader, _ := ioutil.ReadAll(w.Body)
	rgx := regexp.MustCompile("<title>(.*?)</title>")
	if rgx.FindString(string(bodyReader)) != "" {
		assert.Equal(t, rgx.FindStringSubmatch(string(bodyReader))[1], "Title")
	}
	assert.Equal(t, w.Code, 200)
}

func UserLoginMissingParams(t *testing.T) {

	user := models.User{Password: []byte("123456")}
	if err := faker.FakeData(&user); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	if err := user.Create(); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	w := tests.PerformRequest(
		router,
		"POST",
		"/user/login",
		strings.NewReader(
			routers.CreateDataParams(map[string]string{
				"username": user.Username,
			}),
		),
	)

	assert.Equal(t, w.Code, 400)
}

func UserLoginWrongCredentials(t *testing.T) {
	w := tests.PerformRequest(
		router,
		"POST",
		"/user/login",
		strings.NewReader(
			routers.CreateDataParams(map[string]string{
				"username": faker.Username(),
				"password": faker.Password(),
			}),
		),
	)

	assert.Equal(t, w.Code, 400)
}

func UserRegisterSuccess(t *testing.T) {
	// TEST GET REQUEST
	w := tests.PerformRequest(
		router,
		"GET",
		"/user/register",
		nil,
	)
	assert.Equal(t, w.Code, 200)

	w = tests.PerformRequest(
		router,
		"POST",
		"/user/register",
		strings.NewReader(
			routers.CreateDataParams(map[string]string{
				"username":   faker.Username(),
				"password":   "123456",
				"first_name": faker.FirstName(),
				"last_name":  faker.LastName(),
			}),
		),
	)

	bodyReader, _ := ioutil.ReadAll(w.Body)
	rgx := regexp.MustCompile("<title>(.*?)</title>")
	if rgx.FindString(string(bodyReader)) != "" {
		assert.Equal(t, rgx.FindStringSubmatch(string(bodyReader))[1], "Title")
	}
	assert.Equal(t, w.Code, 201)
}

func UserRegisterMissingFields(t *testing.T) {
	w := tests.PerformRequest(
		router,
		"POST",
		"/user/register",
		strings.NewReader(
			routers.CreateDataParams(map[string]string{
				"username":   faker.Username(),
				"first_name": faker.FirstName(),
				"last_name":  faker.LastName(),
			}),
		),
	)

	bodyReader, _ := ioutil.ReadAll(w.Body)
	rgx := regexp.MustCompile("<title>(.*?)</title>")
	if rgx.FindString(string(bodyReader)) != "" {
		assert.Equal(t, rgx.FindStringSubmatch(string(bodyReader))[1], "Title")
	}
	assert.Equal(t, w.Code, 400)
}

func UserRegisterUsernameExists(t *testing.T) {
	username := faker.Username()
	user := models.User{Password: []byte("123456"), Username: username}
	if err := user.Create(); err != nil {
		t.Fail()
	}
	w := tests.PerformRequest(
		router,
		"POST",
		"/user/register",
		strings.NewReader(
			routers.CreateDataParams(map[string]string{
				"username":   username,
				"password":   faker.Password(),
				"first_name": faker.FirstName(),
				"last_name":  faker.LastName(),
			}),
		),
	)

	bodyReader, _ := ioutil.ReadAll(w.Body)
	rgx := regexp.MustCompile("<title>(.*?)</title>")
	if rgx.FindString(string(bodyReader)) != "" {
		assert.Equal(t, rgx.FindStringSubmatch(string(bodyReader))[1], "Title")
	}
	assert.Equal(t, w.Code, 400)
}

func UserMe(t *testing.T) {
	user := models.User{Password: []byte("123456"), IsActive: true}
	if err := faker.FakeData(&user); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	if err := user.Create(); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	w := tests.PerformRequest(
		router,
		"POST",
		"/user/login",
		strings.NewReader(
			routers.CreateDataParams(map[string]string{
				"username": user.Username,
				"password": "123456",
			}),
		),
	)
	w = tests.PerformRequest(
		router,
		"GET",
		"/user/me?token="+w.Result().Cookies()[0].Value,
		nil,
	)

	bodyReader, _ := ioutil.ReadAll(w.Body)
	fmt.Println(string(bodyReader))
	assert.Equal(t, w.Code, 200)
}
