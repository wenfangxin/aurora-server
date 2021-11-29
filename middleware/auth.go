package middleware

import (
	"aurora/core"
	"aurora/service"
	"github.com/gin-gonic/gin"
	"strings"
)

var srv = service.New()

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		//token := ""
		url := UrlSplit(ctx.Request.RequestURI)
		if token == "" {
			ctx.JSON(200, gin.H{
				"code": 403,
				"msg":  "缺少令牌验证",
			})
			ctx.Abort()
			return
		}

		claims, err := core.ParseToken(token)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": 403,
				"msg":  "令牌不正确或已过期",
			})
			ctx.Abort()
			return
		}

		if claims == nil {
			ctx.JSON(200, gin.H{
				"code": 403,
				"msg":  "参数解析失败",
			})
			ctx.Abort()
			return
		}

		user, err := srv.GetUserData(claims.UserId)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": 403,
				"msg":  "用户信息获取失败",
			})
			ctx.Abort()
			return
		}

		if user.State == 0 {
			ctx.JSON(200, gin.H{
				"code": 403,
				"msg":  "账号已被封禁",
			})
			ctx.Abort()
			return
		}

		apiList, err := srv.GetRoleApiList(claims.RoleKey)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": 403,
				"msg":  "获取api权限失败",
			})
			ctx.Abort()
			return
		}
		_, isExist := Find(apiList, url)

		if !isExist {
			ctx.JSON(200, gin.H{
				"code": 401,
				"msg":  "暂无权限",
			})
			ctx.Abort()
			return
		}

		ctx.Keys = map[string]interface{}{
			"userId":  claims.UserId,
			"roleKey": claims.RoleKey,
		}

		ctx.Next()

	}
}

func UrlSplit(url string) string {
	path := strings.Split(url, "?")
	return path[0]
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
