package middlewares

import (
	"github.com/appleboy/gin-jwt/v2"
	"sancap/internal/models"
	"time"
)

func Authentication() *jwt.GinJWTMiddleware {
	authMiddleware, authErr := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "test",
		Key: []byte("somesting"),
		Timeout: time.Hour,
		MaxRefresh: time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.UserModel)
		},
	})
}