package utils

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

func SetJWTCookie(token string, c *gin.Context, mw *jwt.GinJWTMiddleware) {
	c.SetCookie(
		mw.CookieName,
		token,
		int(mw.TimeFunc().Add(mw.Timeout).Unix()-time.Now().Unix()),
		"/",
		mw.CookieDomain,
		mw.SecureCookie,
		mw.CookieHTTPOnly,
	)
}
