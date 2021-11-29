package models

type MenuRouter struct {
	Id          int           `json:"id" gorm:"->"`
	OrderNumber int           `json:"order_number"`
	Path        string        `json:"path" binding:"required"`
	Name        string        `json:"name" binding:"required"`
	Redirect    string        `json:"redirect"`
	Component   string        `json:"component" binding:"required"`
	ParentId    int           `json:"parent_id"`
	Title       string        `json:"title" binding:"required"`
	EnTitle     string        `json:"en_title" binding:"required"`
	Icon        string        `json:"icon"`
	Keep        int           `json:"keep"`
	Affix       int           `json:"affix"`
	Hidden      int           `json:"hidden"`
	Nested      int           `json:"nested"`
	Children    *[]MenuRouter `json:"children,omitempty" gorm:"foreignKey:ParentId"`
	CreateTime  *TimeNormal   `gorm:"->" json:"create_time"`
	UpdateTime  *TimeNormal   `gorm:"->" json:"update_time"`
}

type Role struct {
	Id          int         `json:"id" gorm:"->"`
	RoleKey     string      `json:"role_key" binding:"required"`
	RoleName    string      `json:"role_name" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Menu        string      `json:"menu" binding:"required"`
	Api         string      `json:"api" binding:"required"`
	CreateTime  *TimeNormal `gorm:"->" json:"create_time"`
	UpdateTime  *TimeNormal `gorm:"->" json:"update_time"`
}

type RoleArg struct {
	RoleKey  string `json:"role_key" binding:"required"`
	RoleName string `json:"role_name" binding:"required"`
}

type User struct {
	Id         int         `json:"id" gorm:"->"`
	Avatar     string      `json:"avatar"`
	UserName   string      `json:"user_name" binding:"required"`
	Number     string      `json:"number" binding:"required"`
	Phone      string      `json:"phone" binding:"required"`
	Password   string      `json:"-"`
	RoleKey    string      `json:"role_key" binding:"required"`
	RoleData   *RoleArg    `json:"role_data,omitempty" gorm:"foreignKey:RoleKey;references:RoleKey"`
	Sex        int         `json:"sex" binding:"required"`
	State      int         `json:"state" gorm:"->"`
	CreateTime *TimeNormal `gorm:"->" json:"create_time"`
}

type QueryUserArg struct {
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	UserName string `json:"user_name"`
	Number   string `json:"number" `
	Phone    string `json:"phone"`
	RoleKey  string `json:"role_key"`
	State    *int   `json:"state"`
}

type UpdateUserStateArg struct {
	State *int `json:"state" binding:"required"`
	Id    int  `json:"id" binding:"required"`
}

func (r *RoleArg) TableName() string {
	return "roles"
}

type Login struct {
	Number   string `json:"number" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ModuleData struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	EnTitle string `json:"en_title"`
}

func (m *ModuleData) TableName() string {
	return "menu_router"
}

type Api struct {
	Id         int         `json:"id" gorm:"->"`
	Name       string      `json:"name"`
	Url        string      `json:"url"`
	Method     string      `json:"method"`
	Module     int         `json:"module"`
	ModuleData *ModuleData `json:"module_data" gorm:"foreignKey:Module"`
	CreateTime TimeNormal  `json:"create_time" gorm:"->"`
}
