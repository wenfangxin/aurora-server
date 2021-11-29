package config

type MysqlConfig struct {
	Name     string
	Password string
	Host     string
	Port     string
	DBName   string
}

type WxAppConfig struct {
	AddId     string
	Secret    string
	GrantType string
}

type EmailConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type QINiuConfig struct {
	AccessKey string
	SecretKey string
	Bucket string
}

var JwtSinge = "XXXXXXXXX" //自定义

func Config() *MysqlConfig {
	config := &MysqlConfig{
		Name:     "root",
		Password: "xxxx",
		Host:     "xxx",
		Port:     "3306",
		DBName:   "xxx",
	}
	return config
}

func WxApp() *WxAppConfig {
	config := &WxAppConfig{
		AddId:     "xx",
		Secret:    "xx",
		GrantType: "xx",
	}
	return config
}

func Email() *EmailConfig {
	config := &EmailConfig{
		Host:     "smtp.163.com",
		Port:     465,
		User:     "xxxxx@163.com",
		Password: "xxxxx",
	}
	return config
}

func QiNiu() *QINiuConfig {
	config := QINiuConfig{
		AccessKey: "xxxx",
		SecretKey: "xxxx-xxxx",
		Bucket: "xxxxx",
	}
	return &config
}
