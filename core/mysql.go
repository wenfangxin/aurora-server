package core

import (
	"aurora/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitMysql() *gorm.DB {
	conf := config.Config()
	//数据库连接信息 parseTime解析时间 loc本地时间
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Name, conf.Password, conf.Host, conf.Port, conf.DBName)
	//sql日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
	//连接数据库
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableAutomaticPing: false,     //GORM 会自动 ping 数据库以检查数据库的可用性，若要禁用该特性，可将其设置为 true
		PrepareStmt:          true,      //所有的操作都会创建并缓存预编译语句，以加速后续执行速度
		Logger:               newLogger, //sql日志设置
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
		},
	})

	if err != nil {
		panic(err)
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, sqlDBErr := DB.DB()
	if sqlDBErr != nil {
		panic(sqlDBErr)
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return DB
}
