package web

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"sancap/internal/handlers"
	"sancap/internal/models"
	"sancap/tests"
	"strings"
	"testing"
)

type testFunc func(t *testing.T, router *gin.Engine, handler handlers.BaseHandler)

func TestUser(t *testing.T) {
	tests.SetupTest()

	funcs := []testFunc{
		UserLoginTest,
		UserLoginMissingParams,
		UserLoginWrongCredentials,
		UserRegisterSuccess,
		UserRegisterMissingFields,
		UserRegisterUsernameExists,
		UserMeTest,
		UserChangePasswordFailure,
		UserChangePasswordSuccess,
	}
	for _, f := range funcs {
		f := f
		funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		funcName = filepath.Ext(funcName)
		funcName = strings.TrimPrefix(funcName, ".")
		t.Run(funcName, func(t *testing.T) {
			t.Parallel()
			h, r := tests.SetupTestRouter()
			f(t, r, h)
		})
	}
}

func UserLoginTest(t *testing.T, router *gin.Engine, handler handlers.BaseHandler) {
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

	if !user.Create(handler.DB) {
		t.Fail()
	}
	w = tests.PerformRequest(
		router,
		"POST",
		"/user/login",
		strings.NewReader(
			tests.CreateTestParams(map[string]string{
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
	handler.DB.Close()
}

func UserLoginMissingParams(t *testing.T, router *gin.Engine, handler handlers.BaseHandler) {
	user := models.User{Password: []byte("123456")}
	if err := faker.FakeData(&user); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	if !user.Create(handler.DB) {
		t.Fail()
	}
	w := tests.PerformRequest(
		router,
		"POST",
		"/user/login",
		strings.NewReader(
			tests.CreateTestParams(map[string]string{
				"username": user.Username,
			}),
		),
	)

	assert.Equal(t, w.Code, 401)
}

func UserLoginWrongCredentials(t *testing.T, router *gin.Engine, handler handlers.BaseHandler) {
	w := tests.PerformRequest(
		router,
		"POST",
		"/user/login",
		strings.NewReader(
			tests.CreateTestParams(map[string]string{
				"username": faker.Username(),
				"password": faker.Password(),
			}),
		),
	)

	assert.Equal(t, w.Code, 401)
}

func UserRegisterSuccess(t *testing.T, router *gin.Engine, handler handlers.BaseHandler) {
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
			tests.CreateTestParams(map[string]string{
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

func UserRegisterMissingFields(t *testing.T, router *gin.Engine, handler handlers.BaseHandler) {
	w := tests.PerformRequest(
		router,
		"POST",
		"/user/register",
		strings.NewReader(
			tests.CreateTestParams(map[string]string{
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

func UserRegisterUsernameExists(t *testing.T, router *gin.Engine, handler handlers.BaseHandler) {
	username := faker.Username()
	user := models.User{Password: []byte("123456"), Username: username}
	if !user.Create(handler.DB) {
		t.Fail()
	}
	w := tests.PerformRequest(
		router,
		"POST",
		"/user/register",
		strings.NewReader(
			tests.CreateTestParams(map[string]string{
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

func UserMeTest(t *testing.T, router *gin.Engine, handler handlers.BaseHandler) {
	user := models.User{Password: []byte("123456"), IsActive: true}
	if err := faker.FakeData(&user); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	if !user.Create(handler.DB) {
		t.Fail()
	}

	w := tests.PerformRequest(
		router,
		"POST",
		"/user/login",
		strings.NewReader(
			tests.CreateTestParams(map[string]string{
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

func UserChangePasswordSuccess(t *testing.T, router *gin.Engine, handler handlers.BaseHandler) {
	user := models.User{Password: []byte("123456"), IsActive: true}
	if err := faker.FakeData(&user); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	if !user.Create(handler.DB) {
		t.Fail()
	}

	w := tests.PerformRequest(
		router,
		"POST",
		"/user/login",
		strings.NewReader(
			tests.CreateTestParams(map[string]string{
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
			tests.CreateTestParams(map[string]string{
				"old_password":  "123456",
				"new_password":  "123123",
				"new_password2": "123123",
			}),
		),
	)

	assert.Equal(t, w.Code, 200)
}

func UserChangePasswordFailure(t *testing.T, router *gin.Engine, handler handlers.BaseHandler) {
	user := models.User{Password: []byte("123456"), IsActive: true}
	if err := faker.FakeData(&user); err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	if !user.Create(handler.DB) {
		t.Fail()
	}

	w := tests.PerformRequest(
		router,
		"POST",
		"/user/login",
		strings.NewReader(
			tests.CreateTestParams(map[string]string{
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
			tests.CreateTestParams(map[string]string{
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
			tests.CreateTestParams(map[string]string{
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
			tests.CreateTestParams(map[string]string{
				"old_password":  "123456",
				"new_password":  "1231234",
				"new_password2": "123123",
			}),
		),
	)

	assert.Equal(t, w.Code, 400)
}
