package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type ResponseCode int

const SuccessCode ResponseCode = iota

type Response struct {
	ErrorCode ResponseCode `json:"error_code"`
	ErrorMsg  string       `json:"error_msg"`
	Data      interface{}  `json:"data"`
}

// TODO: implement with api
func ResponseError(ctx *gin.Context, code ResponseCode, err error) {
	resp := &Response{
		ErrorCode: code,
		ErrorMsg:  err.Error(),
		Data:      "",
	}
	ctx.JSON(200, resp)
	response, _ := json.Marshal(resp)
	ctx.Set("response", string(response))
	ctx.AbortWithError(200, err)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	resp := &Response{
		ErrorCode: SuccessCode,
		ErrorMsg:  "",
		Data:      data,
	}
	ctx.JSON(200, resp)
	response, _ := json.Marshal(resp)
	ctx.Set("response", string(response))
}
