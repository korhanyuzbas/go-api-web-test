package user

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"regexp"
	"sancap/internal/configs"
	"sancap/internal/models"
	"sancap/internal/routers"
	"sancap/tests"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("blah blah start")
	tests.InitTestConfig()
	code := m.Run()
	fmt.Println("stoppp blahhh")
	configs.DB.DropTableIfExists(&models.User{}, &models.UserVerification{})
	os.Exit(code)
}

func TestUserLogin(t *testing.T) {
	router := tests.SetupTestRouter()

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

func TestUserLogin_MissingParams(t *testing.T) {
	router := tests.SetupTestRouter()

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

func TestUserLogin_WrongCredentials(t *testing.T) {
	router := tests.SetupTestRouter()

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

func TestUserRegister_Success(t *testing.T) {
	router := tests.SetupTestRouter()

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

func TestUserRegister_MissingFields(t *testing.T) {
	router := tests.SetupTestRouter()

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

func TestUserRegister_UsernameExists(t *testing.T) {
	router := tests.SetupTestRouter()
	username := faker.Username()
	user := models.User{Password: []byte("123456"), Username: username}
	user.Create()
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

func TestUserMe(t *testing.T) {
	router := tests.SetupTestRouter()

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
