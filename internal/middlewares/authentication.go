package middlewares

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"sancap/internal/models"
	"time"
)

func Authentication() *jwt.GinJWTMiddleware {
	authMiddleware, authErr := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test",
		Key:         []byte("somesting"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.UserModel); ok {
				return jwt.MapClaims{
					"id": v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.UserModel{
				Username: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.UserModel
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			if err := models.GetUserCred(&loginVals, username, password); err != nil {
				return nil, err
			}
			return &loginVals, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			var authModel models.UserModel
			if v, ok := data.(*models.UserModel); ok {
				if err := models.GetUserByName(&authModel, v.Username); err != nil {
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
