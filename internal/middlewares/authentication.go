package middlewares

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sancap/internal/configs"
	"sancap/internal/dto"
	"sancap/internal/handlers"
	"sancap/internal/helpers"
	"sancap/internal/models"
	"time"
)

func getDefaultMiddleware(cookie bool, handler handlers.BaseHandler) *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:       "test",
		Key:         []byte(configs.AppConfig.SecretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: configs.JwtKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					configs.JwtKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Username: claims[configs.JwtKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			userLoginInput := dto.LoginUserInput{}
			if err := userLoginInput.ShouldBind(c); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := userLoginInput.Username
			password := userLoginInput.Password

			user := models.User{
				Username: username,
				Password: []byte(password),
			}

			if err := user.GetCredentials(handler.DB, userLoginInput.Password); err != nil {
				return nil, err
			}
			return &user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			var authModel models.User
			if v, ok := data.(*models.User); ok {
				if err := authModel.GetByName(handler.DB, v.Username); err != nil {
					return false
				}
				if !authModel.IsActive {
					return false
				}
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		SendCookie:    cookie,
	}
}

func AuthenticationWeb(handler handlers.BaseHandler) *jwt.GinJWTMiddleware {
	defaultMiddleware := getDefaultMiddleware(true, handler)
	defaultMiddleware.LoginResponse = func(context *gin.Context, i int, s string, t time.Time) {
		helpers.SetJWTCookie(s, context, defaultMiddleware)
		context.HTML(http.StatusOK, "user.login.html", gin.H{"error": nil, "user": nil})
	}
	defaultMiddleware.Unauthorized = func(context *gin.Context, i int, s string) {
		context.HTML(i, "user.login.html", gin.H{"error": s, "user": nil})
	}

	authMiddleware, authErr := jwt.New(defaultMiddleware)
	if authErr != nil {
		log.Fatal("JWT error: " + authErr.Error())
	}

	return authMiddleware
}

func AuthenticationAPI(handler handlers.BaseHandler) *jwt.GinJWTMiddleware {
	defaultMiddleware := getDefaultMiddleware(false, handler)

	authMiddleware, authErr := jwt.New(defaultMiddleware)
	if authErr != nil {
		log.Fatal("JWT error: " + authErr.Error())
	}

	return authMiddleware
}
