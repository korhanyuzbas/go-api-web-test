package user

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"sancap/internal/models"
	"sancap/internal/routers"
	"sancap/tests"
	"strings"
	"testing"
)

var router *gin.Engine

type testFunc func(t *testing.T)

func TestUser(t *testing.T) {
	tests.Setup()
	router = tests.SetupTestRouter()
	funcs := []testFunc{
		UserLogin,
		UserLoginMissingParams,
		UserLoginWrongCredentials,
		UserRegisterSuccess,
		UserRegisterMissingFields,
		UserRegisterUsernameExists,
		UserMe,
		UserChangePasswordFailure,
		UserChangePasswordSuccess,
	}
	for _, f := range funcs {
		f := f
		funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		funcName = filepath.Ext(funcName)
		funcName = strings.TrimPrefix(funcName, ".")
		t.Run(funcName, func(t *testing.T) {
			f(t)
		})
	}
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

	assert.Equal(t, w.Code, 200)
}

func UserChangePasswordSuccess(t *testing.T) {
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
	jwtToken := w.Result().Cookies()[0].Value
	w = tests.PerformRequest(
		router,
		"GET",
		"/user/change_password?token="+jwtToken,
		nil,
	)

	assert.Equal(t, w.Code, 200)

	w = tests.PerformRequest(
		router,
		"POST",
		"/user/change_password?token="+jwtToken,
		strings.NewReader(
			routers.CreateDataParams(map[string]string{
				"old_password":  "123456",
				"new_password":  "123123",
				"new_password2": "123123",
			}),
		),
	)

	assert.Equal(t, w.Code, 200)
}

func UserChangePasswordFailure(t *testing.T) {
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
	jwtToken := w.Result().Cookies()[0].Value

	// wrong old password
	w = tests.PerformRequest(
		router,
		"POST",
		"/user/change_password?token="+jwtToken,
		strings.NewReader(
			routers.CreateDataParams(map[string]string{
				"old_password":  "qweasd",
				"new_password":  "123123",
				"new_password2": "123123",
			}),
		),
	)

	assert.Equal(t, w.Code, 400)

	// missing parameters
	w = tests.PerformRequest(
		router,
		"POST",
		"/user/change_password?token="+jwtToken,
		strings.NewReader(
			routers.CreateDataParams(map[string]string{
				"old_password": "123456",
				"new_password": "123123",
			}),
		),
	)

	// new password error
	w = tests.PerformRequest(
		router,
		"POST",
		"/user/change_password?token="+jwtToken,
		strings.NewReader(
			routers.CreateDataParams(map[string]string{
				"old_password":  "123456",
				"new_password":  "1231234",
				"new_password2": "123123",
			}),
		),
	)

	assert.Equal(t, w.Code, 400)
}
