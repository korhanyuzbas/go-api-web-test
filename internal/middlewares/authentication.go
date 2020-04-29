package middlewares

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"sancap/internal/configs"
	"sancap/internal/models"
	"time"
)

func Authentication() *jwt.GinJWTMiddleware {
	authMiddleware, authErr := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test",
		Key:         []byte("somesting"),
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
			var user models.User
			if err := c.ShouldBind(&user); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := user.Username
			password := user.Password

			if err := user.GetCredentials(username, password); err != nil {
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
	})
	if authErr != nil {
		log.Fatal("JWT error: " + authErr.Error())
	}
	return authMiddleware
}
