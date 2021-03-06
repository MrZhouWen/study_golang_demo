package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code": 0, //0表示成功,其他表示失败
	"message":"success"， //用来描述失败的原因
	"data":{

	}
}
*/
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(ctx *gin.Context, code int) {

	responseData := &ResponseData{
		Code:    1,
		Message: "失败",
	}

	ctx.JSON(http.StatusOK, responseData)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	responseData := &ResponseData{
		Code:    0,
		Message: "成功",
		Data:    data,
	}

	ctx.JSON(http.StatusOK, responseData)
}
