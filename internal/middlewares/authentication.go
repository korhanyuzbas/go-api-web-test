package middlewares

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"sancap/internal/configs"
	"sancap/internal/dto"
	"sancap/internal/models"
	"time"
)

func getDefaultMiddleware() *jwt.GinJWTMiddleware {
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

			if err := user.GetCredentials(userLoginInput.Password); err != nil {
				return nil, err
			}
			return &user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			var authModel models.User
			if v, ok := data.(*models.User); ok {
				if err := authModel.GetByName(v.Username); err != nil {
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
		SendCookie:    true,
	}
}

func AuthenticationWeb() *jwt.GinJWTMiddleware {
	defaultMiddleware := getDefaultMiddleware()
	defaultMiddleware.SendCookie = true

	authMiddleware, authErr := jwt.New(defaultMiddleware)
	if authErr != nil {
		log.Fatal("JWT error: " + authErr.Error())
	}

	return authMiddleware
}

func AuthenticationAPI() *jwt.GinJWTMiddleware {
	defaultMiddleware := getDefaultMiddleware()
	defaultMiddleware.SendCookie = false

	authMiddleware, authErr := jwt.New(defaultMiddleware)
	if authErr != nil {
		log.Fatal("JWT error: " + authErr.Error())
	}

	return authMiddleware
}
