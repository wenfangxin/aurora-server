package http

import (
	"aurora/middleware"
	"aurora/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var srv *service.Service

func Init() {
	srv = service.New()

	//web 引擎
	app := gin.Default()

	//中间件
	app.Use(middleware.ExceptionRecover()) //全局异常处理
	app.Use(middleware.Cors())             //跨域处理
	//app.Use(middleware.Log())         //开启日志

	//路由
	router(app)

	if err := app.Run(":3000"); err != nil {
		panic(err)
	}

}

func router(app *gin.Engine) {
	{
		//静态html文件
		app.StaticFS("/html", http.Dir("./html"))
		//静态资源访问路径
		app.StaticFS("/static", http.Dir("./static"))

		app.POST("/login",Login)
	}

	//以下都是需要进行权限校验的接口
	app.Use(middleware.Auth())

	//用户相关
	permissions := app.Group("/permissions")
	{
		permissions.POST("/upload", Upload)
		permissions.GET("/getQiNiuToken", GetQiNiuToken)
		permissions.GET("/getMenuRouter", GetMenuRouter)
		permissions.GET("/getMenuList", GetMenuList)
		permissions.POST("/addMenuRouter", AddMenuRouter)
		permissions.GET("/GetSingleMenuRouter", GetSingleMenuRouter)
		permissions.POST("/updateMenuRouter", UpdateMenuRouter)
		permissions.GET("/deleteMenuRouter", DeleteMenuRouter)
		permissions.POST("/createRole", CreateRole)
		permissions.GET("/getRolesList", GetRolesList)
		permissions.POST("/updateRole", UpdateRole)
		permissions.GET("/deleteRole", DeleteRole)
		permissions.POST("/createUser", CreateUser)
		permissions.GET("/getUserList", GetUserList)
		permissions.POST("/updateUser", UpdateUser)
		permissions.POST("/updateUserState", UpdateUserSate)
		permissions.GET("/deleteUser", DeleteUser)
		permissions.POST("/queryUser", QueryUser)
		permissions.GET("/getUserData",GetUserData)
		permissions.GET("/getApiList",GetApiList)
		permissions.POST("/createApi",CreateApi)
		permissions.POST("/deleteApi",DeleteApi)
		permissions.POST("/updateApi",UpdateApi)
	}

}
