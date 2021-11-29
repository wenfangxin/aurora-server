package http

import (
	"aurora/core"
	"aurora/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

// GetMenuRouter 获取用户菜单
func GetMenuRouter(ctx *gin.Context) {
	roleKey := ctx.Keys["roleKey"]
	result, err := srv.MenuRouter(roleKey.(string))
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, result)
}

// GetMenuList 获取菜单列表
func GetMenuList(ctx *gin.Context) {
	result, err := srv.GetMenuList()
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, result)
}

// AddMenuRouter 新增menuRouter
func AddMenuRouter(ctx *gin.Context) {
	var menuRouter models.MenuRouter
	if err := ctx.ShouldBind(&menuRouter); err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	if err := srv.CreateMenuRouter(&menuRouter); err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, nil)
}

// GetSingleMenuRouter 根据id获取menu
func GetSingleMenuRouter(ctx *gin.Context) {
	id, exists := ctx.GetQuery("id")
	menuId, _ := strconv.Atoi(id)
	if !exists {
		core.ErrorMsg(ctx, "缺少参数")
		return
	}
	result, err := srv.GetSingleMenuData(menuId)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}

	core.SuccessMsg(ctx, result)
}

// UpdateMenuRouter 修改数据菜单
func UpdateMenuRouter(ctx *gin.Context) {
	menu := new(models.MenuRouter)
	err := ctx.ShouldBind(menu)
	if menu.Id == 0 {
		core.ErrorMsg(ctx, "缺少Id")
		return
	}
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	err = srv.UpdateMenuRouter(menu)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, nil)

}

// DeleteMenuRouter 删除menuRouter
func DeleteMenuRouter(ctx *gin.Context) {
	id, exists := ctx.GetQuery("id")
	menuId, _ := strconv.Atoi(id)
	if !exists {
		core.ErrorMsg(ctx, "缺少参数")
		return
	}
	err := srv.DeleteMenuRouter(menuId)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, nil)
}

// CreateRole 创建角色
func CreateRole(ctx *gin.Context) {
	role := new(models.Role)
	err := ctx.ShouldBind(role)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	err = srv.CreateRole(role)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, nil)
}

//GetRolesList 获取所有角色列表
func GetRolesList(ctx *gin.Context) {
	var paging = new(models.Paging)
	err := ctx.ShouldBind(paging)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	list, err := srv.GetRolesList(paging)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, list)
}

// UpdateRole 更新角色
func UpdateRole(ctx *gin.Context) {
	role := new(models.Role)
	err := ctx.ShouldBind(role)
	if role.Id == 0 {
		core.ErrorMsg(ctx, "缺少Id")
		return
	}
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	err = srv.UpdateRole(role)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, nil)
}

// DeleteRole 删除角色
func DeleteRole(ctx *gin.Context) {
	id, exists := ctx.GetQuery("id")
	idInt, _ := strconv.Atoi(id)
	if !exists {
		core.ErrorMsg(ctx, "缺少参数ID")
		return
	}
	err := srv.DeleteRole(idInt)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, nil)

}

// CreateUser 创建用户
func CreateUser(ctx *gin.Context) {
	user := new(models.User)
	err := ctx.ShouldBind(user)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	err = srv.CreateUser(user)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, nil)

}

// GetUserList 获取用户列表
func GetUserList(ctx *gin.Context) {
	paging := new(models.Paging)
	err := ctx.ShouldBind(paging)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	list, err := srv.GetUserList(paging)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}

	core.SuccessMsg(ctx, list)
}

// UpdateUser 更新用户信息
func UpdateUser(ctx *gin.Context) {
	user := new(models.User)
	err := ctx.ShouldBind(user)
	if user.Id == 0 {
		core.ErrorMsg(ctx, "缺少Id")
		return
	}
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	err = srv.UpdateUser(user)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, nil)

}

// UpdateUserSate 更新用户信息
func UpdateUserSate(ctx *gin.Context) {
	arg := new(models.UpdateUserStateArg)
	err := ctx.ShouldBind(arg)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	err = srv.UpdateUserState(arg)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, nil)
}

//DeleteUser 删除用户
func DeleteUser(ctx *gin.Context) {
	id, exists := ctx.GetQuery("id")
	idInt, _ := strconv.Atoi(id)
	fmt.Println(id)
	if !exists {
		core.ErrorMsg(ctx, "缺少参数Id")
		return
	}
	err := srv.DeleteUser(idInt)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, nil)
}

// QueryUser 查询用户
func QueryUser(ctx *gin.Context) {
	userQuery := new(models.QueryUserArg)
	err := ctx.ShouldBind(userQuery)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	user, err := srv.QueryUser(userQuery)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx, user)
}

//Login 用户登录
func Login(ctx *gin.Context) {
	loginData := new(models.Login)
	err := ctx.ShouldBind(loginData)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	user, err := srv.UserIsExists(loginData)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	if user["id"] == nil {
		core.ErrorMsg(ctx, "账号不存在")
		return
	}

	if loginData.Password != user["password"] {
		core.ErrorMsg(ctx, "密码错误")
		return
	}
	if user["state"] == int32(0) {
		core.ErrorMsg(ctx, "账号已被封禁")
		return
	}

	token, err := core.CreateJwtToken(user["id"].(int32), user["role_key"].(string))
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}

	core.SuccessMsg(ctx, gin.H{
		"token":    token,
		"user_id":  user["id"],
		"role_key": user["role_key"],
	})

}

// GetUserData 获取当前登录用户的个人信息
func GetUserData(ctx *gin.Context) {
	userId := ctx.Keys["userId"]
	data, err := srv.GetUserData(userId.(int32))
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx,data)
}


func GetApiList(ctx *gin.Context)  {
	paging := new(models.Paging)
	err := ctx.ShouldBind(paging)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	list, err := srv.GetApiList(paging)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx,list)
}

func CreateApi (ctx *gin.Context){
	api := new(models.Api)
	err := ctx.ShouldBind(api)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	err = srv.CreateApi(api)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx,nil)
}


func DeleteApi(ctx *gin.Context)  {
	data,err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	err = srv.DeleteApi(data)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx,nil)
}

func UpdateApi(ctx *gin.Context)  {
	api := new(models.Api)
	err := ctx.ShouldBind(api)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	err = srv.UpdateApi(api)
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	core.SuccessMsg(ctx,nil)
}