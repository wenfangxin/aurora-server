package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func ExceptionRecover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err:= recover();err!=nil{
				color := string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
				reset := string([]byte{27, 91, 48, 109})
				//打印错误堆栈信息
				log.Println(color, err, reset)
				//记录日志
				Logger().WithFields(logrus.Fields{}).Error(err)
				ctx.JSON(http.StatusOK,gin.H{
					"code":500,
					"msg":errorToString(err),
				})
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				ctx.Abort()
			}
		}()
		//加载完 defer recover，继续后续接口调用
		ctx.Next()
	}
}


// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
