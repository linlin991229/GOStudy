package models

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"lin.com/im/utils"
)

var DB *gorm.DB

func init() {
	// 读取配置文件
	cfg, errIni := ini.Load("./conf/app.ini")

	if errIni != nil {
		log.Panicln(errIni)
	}
	// 获取MySQL配置
	mysqlConf := cfg.Section("mysql")

	username := mysqlConf.Key("username").String()
	password := mysqlConf.Key("password").String()
	host := mysqlConf.Key("host").String()
	port := mysqlConf.Key("port").String()
	database := mysqlConf.Key("database").String()

	// 连接数据库
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	dbLogger := logger.New(
		utils.LogInfo,
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 彩色打印
		},
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dbLogger})
	if err != nil {
		log.Panicln(err)
	}
}
