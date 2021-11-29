package core

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ErrorMsg 错误 返回
func ErrorMsg(ctx *gin.Context,msg interface{}) {
	color := string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	reset := string([]byte{27, 91, 48, 109})
	//打印错误堆栈信息
	log.Println(color, msg, reset)
	ctx.JSON(http.StatusOK,gin.H{
		"code":500,
		"msg":msg,
	})
}

// SuccessMsg 成功 返回
func SuccessMsg (ctx *gin.Context,data interface{}) {
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":data,
	})
}
