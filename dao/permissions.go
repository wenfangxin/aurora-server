package dao

import (
	"aurora/models"
	"fmt"
	"gorm.io/gorm"
)

var roleID []int

// RecursionFunction 嵌套路由递归查找
func (d *Dao) RecursionFunction(db *gorm.DB) *gorm.DB {
	return db.Preload("Children", d.RecursionFunction).Order("order_number desc").Where("menu_router.state = ? AND menu_router.id IN ?", 1, roleID)
}

// SelectUserMenuRouter 获取当前用户的菜单路由
func (d *Dao) SelectUserMenuRouter(roleId []int) (*[]models.MenuRouter, error) {
	roleID = roleId
	menuRouterList := new([]models.MenuRouter)
	err := d.DB.Table("menu_router").Where(
		"menu_router.state = ? AND menu_router.parent_id = ? AND id IN ?", 1, 0, roleID).Order(
		"order_number desc").Preload(
		"Children", d.RecursionFunction).Find(menuRouterList).Error
	return menuRouterList, err
}

func (d *Dao) MenuRecursionFunction(db *gorm.DB) *gorm.DB {
	return db.Preload("Children", d.MenuRecursionFunction).Order("order_number desc").Where("menu_router.state = ?", 1)
}

// GetMenuList 获取菜单列表
func (d *Dao) GetMenuList() (*[]models.MenuRouter, error) {
	menuRouterList := new([]models.MenuRouter)
	err := d.DB.Table("menu_router").Where(
		"menu_router.state = ? AND menu_router.parent_id = ?", 1, 0).Order(
		"order_number desc").Preload(
		"Children", d.MenuRecursionFunction).Find(menuRouterList).Error
	return menuRouterList, err
}

// CreateMenuRouter 创建菜单
func (d *Dao) CreateMenuRouter(menuRouter *models.MenuRouter) error {
	return d.DB.Table("menu_router").Create(menuRouter).Error
}

// GetSingleMenu 根据menu id 查找 数据
func (d *Dao) GetSingleMenu(id int) (*models.MenuRouter, error) {
	menuRoute := new(models.MenuRouter)
	err := d.DB.Table("menu_router").Where("menu_router.id = ? AND menu_router.state = ?", id, 1).Find(menuRoute).Error
	return menuRoute, err
}

// UpdateMenuRouter 根据菜单id更新菜单
func (d *Dao) UpdateMenuRouter(router *models.MenuRouter) error {
	return d.DB.Table("menu_router").Omit("id", "parent_id").Where("id = ?", router.Id).Save(router).Error
}

// DeleteMenuRouter 删除菜单
func (d *Dao) DeleteMenuRouter(id int) (err error) {
	childIdList := make([]int, 0)
	d.DB.Table("menu_router").Where("parent_id = ?", id).Select("id").Find(&childIdList)
	//删除菜单下的子菜单
	if len(childIdList) != 0 {
		err = d.DB.Debug().Table("menu_router").Where("id IN ?", childIdList).Delete("").Error
	}
	err = d.DB.Table("menu_router").Where("id = ?", id).Delete("").Error

	return
}

// CreateRole 创建角色
func (d *Dao) CreateRole(role *models.Role) error {
	return d.DB.Table("roles").Create(role).Error
}

// GetRolesList 获取所有角色list
func (d *Dao) GetRolesList(paging *models.Paging) (roles *[]models.Role, count int64, err error) {
	err = d.DB.Table("roles").Count(&count).Offset((paging.Page - 1) * paging.Size).Limit(paging.Size).Find(&roles).Error
	return
}

// UpdateRole 更新修改role信息
func (d *Dao) UpdateRole(role *models.Role) (err error) {
	err = d.DB.Table("roles").Where("id = ?", role.Id).Updates(role).Error
	return
}

// DeleteRole 删除role
func (d *Dao) DeleteRole(id int) (err error) {
	err = d.DB.Table("roles").Where("id = ?", id).Delete(nil).Error
	return
}

// GetRole 根据key获取角色
func (d *Dao) GetRole(key string) (role *models.Role, err error) {
	err = d.DB.Table("roles").Where("role_key = ?", key).Find(&role).Error
	return
}

// CreateUser 创建用户
func (d *Dao) CreateUser(user *models.User) error {
	return d.DB.Table("user").Create(user).Error
}

// GetUserList 获取用户列表
func (d *Dao) GetUserList(paging *models.Paging) (user *[]models.User, count int64, err error) {
	err = d.DB.Table("user").Count(&count).Offset((paging.Page - 1) * paging.Size).Limit(paging.Size).Preload("RoleData").Find(&user).Error
	return
}

//UpdateUser 更新用户信息
func (d *Dao) UpdateUser(user *models.User) error {
	return d.DB.Table("user").Where("id = ?", user.Id).Updates(user).Error
}

// DeleteUser 删除用户信息
func (d *Dao) DeleteUser(id int) error {
	return d.DB.Table("user").Where("id = ?", id).Delete("").Error
}

// UpdateUserState 更新用户状态 0 封禁 1 正常
func (d *Dao) UpdateUserState(arg *models.UpdateUserStateArg) error {
	return d.DB.Table("user").Where("id = ?", arg.Id).Update("state", arg.State).Error
}

// QueryUser 查询user
func (d *Dao) QueryUser(arg *models.QueryUserArg) (user *[]models.User, count int64, err error) {
	var sql = fmt.Sprintf("user_name LIKE '%s' and number LIKE '%s' and phone LIKE '%s' and role_key LIKE '%s'",
		"%"+arg.UserName+"%",
		"%"+arg.Number+"%",
		"%"+arg.Phone+"%",
		"%"+arg.RoleKey+"%")
	tx := d.DB.Table("user")
	if *arg.State == 2 {
		tx.Where(sql)
	} else {
		tx.Where(sql).Where("state = ?", arg.State)
	}
	tx.Count(&count).Offset((arg.Page - 1) * arg.Size).Limit(arg.Size).Order("create_time desc").Preload("RoleData").Find(&user)
	err = tx.Error
	return
}

// UserIsExists 根据登录账号查询用户
func (d *Dao) UserIsExists(number string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	err := d.DB.Table("user").Where("number = ?", number).Select("id", "password", "state", "role_key").Find(&data).Error
	return data, err
}

// GetUserData 获取用户个人信息
func (d *Dao) GetUserData(userId int32) (user *models.User, err error) {
	err = d.DB.Table("user").Where("id = ?", userId).Find(&user).Error
	return
}

// GetApiList 获取api list
func (d *Dao) GetApiList(paging *models.Paging) (api *[]models.Api, count int64, err error) {
	err = d.DB.Table("api").Count(&count).Offset((paging.Page - 1) * paging.Size).Limit(paging.Size).Preload("ModuleData").Find(&api).Error
	return
}

// CreateApi 创建api
func (d *Dao) CreateApi(api *models.Api) error {
	return d.DB.Table("api").Create(api).Error
}

// DeleteApi delete api
func (d *Dao) DeleteApi(ids []int) error {
	return d.DB.Table("api").Where("id IN ?", ids).Delete("").Error
}

// UpdateApi 修改APi
func (d *Dao) UpdateApi(api *models.Api) error {
	return d.DB.Table("api").Where("id = ?", api.Id).Updates(api).Error
}

// GetRoleApiIdList get role api idList
func (d *Dao) GetRoleApiIdList(key string) (role *models.Role, err error) {
	err = d.DB.Table("roles").Where("role_key = ?", key).Find(&role).Error
	return
}

// GetApiPathList 获取api path list
func (d *Dao) GetApiPathList(keys []int) ([]string, error) {
	var list = make([]string, 0)
	err := d.DB.Table("api").Where("id IN ?", keys).Select("url").Find(&list).Error
	return list, err
}
