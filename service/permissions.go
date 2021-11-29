package service

import (
	"aurora/models"
	"encoding/json"
	"strconv"
	"strings"
)

// MenuRouter 获取用户菜单
func (s *Service) MenuRouter(roleKey string) (*[]models.MenuRouter, error) {
	role, err := s.dao.GetRole(roleKey)
	//字符串转切片
	roleSplit := strings.Split(role.Menu, ",")
	//字符串切片转int切片
	var intSplit = make([]int, 0)
	for _, v := range roleSplit {
		intId, _ := strconv.Atoi(v)
		intSplit = append(intSplit, intId)
	}
	menuRouterList, err := s.dao.SelectUserMenuRouter(intSplit)
	return menuRouterList, err
}

func (s *Service) GetMenuList() (*[]models.MenuRouter, error) {
	return s.dao.GetMenuList()
}

// CreateMenuRouter 创建菜单
func (s *Service) CreateMenuRouter(menuRouter *models.MenuRouter) error {
	return s.dao.CreateMenuRouter(menuRouter)
}

// GetSingleMenuData 根据菜单id查找单个数据
func (s *Service) GetSingleMenuData(id int) (*models.MenuRouter, error) {
	menu, err := s.dao.GetSingleMenu(id)
	return menu, err
}

// UpdateMenuRouter 根据menu id 修改菜单数据
func (s *Service) UpdateMenuRouter(router *models.MenuRouter) error {
	err := s.dao.UpdateMenuRouter(router)
	return err
}

// DeleteMenuRouter 删除menu
func (s *Service) DeleteMenuRouter(id int) error {
	return s.dao.DeleteMenuRouter(id)
}

// CreateRole 创建角色
func (s *Service) CreateRole(role *models.Role) error {
	return s.dao.CreateRole(role)
}

// GetRolesList 获取角色列表
func (s *Service) GetRolesList(paging *models.Paging) (map[string]interface{}, error) {
	var result = make(map[string]interface{})
	data, count, err := s.dao.GetRolesList(paging)
	result["rows"] = data
	result["total"] = count
	return result, err
}

// UpdateRole 修改role数据
func (s *Service) UpdateRole(role *models.Role) error {
	return s.dao.UpdateRole(role)
}

// DeleteRole 删除角色
func (s *Service) DeleteRole(id int) error {
	return s.dao.DeleteRole(id)
}

// CreateUser 创建用户
func (s *Service) CreateUser(user *models.User) (err error) {
	err = s.dao.CreateUser(user)
	return
}

// GetUserList 获取用户列表
func (s *Service) GetUserList(paging *models.Paging) (map[string]interface{}, error) {
	userList, count, err := s.dao.GetUserList(paging)
	var result = map[string]interface{}{
		"rows":  userList,
		"total": count,
	}
	return result, err
}

// UpdateUser 更新用户
func (s *Service) UpdateUser(user *models.User) error {
	return s.dao.UpdateUser(user)
}

// DeleteUser 删除用户
func (s *Service) DeleteUser(id int) error {
	return s.dao.DeleteUser(id)
}

// UpdateUserState 更新用户状态
func (s *Service) UpdateUserState(arg *models.UpdateUserStateArg) error {
	return s.dao.UpdateUserState(arg)
}

// QueryUser query user
func (s *Service) QueryUser(arg *models.QueryUserArg) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	user, count, err := s.dao.QueryUser(arg)
	result["rows"] = user
	result["total"] = count
	return result, err

}

// UserIsExists 登录时判断用户是否存在
func (s *Service) UserIsExists(login *models.Login) (map[string]interface{}, error) {
	return s.dao.UserIsExists(login.Number)
}

// GetUserData 获取当前登录用户信息
func (s *Service) GetUserData(userId int32) (*models.User, error) {
	return s.dao.GetUserData(userId)
}

// GetApiList  获取apiList
func (s *Service) GetApiList(paging *models.Paging) (map[string]interface{}, error) {
	apiList, count, err := s.dao.GetApiList(paging)

	result := map[string]interface{}{
		"rows":  apiList,
		"total": count,
	}

	return result, err

}

// CreateApi 创建api
func (s *Service) CreateApi(api *models.Api) error {
	return s.dao.CreateApi(api)
}

// DeleteApi 删除api
func (s *Service) DeleteApi(data []byte) (err error) {
	mapData := make(map[string][]int)
	err = json.Unmarshal(data, &mapData)
	err = s.dao.DeleteApi(mapData["ids"])
	return
}

// UpdateApi 更新api数据
func (s *Service) UpdateApi(api *models.Api) error {
	return s.dao.UpdateApi(api)
}

// GetRoleApiList 获取角色的api path
func (s *Service) GetRoleApiList(key string) ([]string, error) {
	role, err := s.dao.GetRoleApiIdList(key)
	if err != nil {
		return nil, err
	}
	var apiIds = strings.Split(role.Api, ",")
	var apiIntList = make([]int, 0)
	for _, v := range apiIds {
		intId, _ := strconv.Atoi(v)
		apiIntList = append(apiIntList, intId)
	}
	list, err := s.dao.GetApiPathList(apiIntList)

	if err != nil {
		return nil, err
	}
	return list, err
}
