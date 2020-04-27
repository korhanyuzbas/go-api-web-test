package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "hello canımmm"})
}

//func CreateUser(ctx *gin.Context) {
//	var user models.User
//	if err := ctx.ShouldBindJSON(&user); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//	} else {
//		err := models.CreateUser(&user)
//		if err != nil {
//			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		} else {
//			ctx.JSON(http.StatusCreated, user)
//		}
//	}
//}
//
//func UserMe(ctx *gin.Context) {
//	var user models.User
//	authid, _ := ctx.Get("id")
//	if err := user.GetByName(authid.(*models.User).Username); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		ctx.Abort()
//	}
//	ctx.JSON(http.StatusOK, user)
//}
