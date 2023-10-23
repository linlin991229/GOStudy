package models

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// 读取配置文件
	cfg, errIni := ini.Load("./conf/app.ini")
	if errIni != nil {
		log.Panicln(errIni)
	}
	username := cfg.Section("mysql").Key("user").String()
	password := cfg.Section("mysql").Key("password").String()
	database := cfg.Section("mysql").Key("db_name").String()
	port := cfg.Section("mysql").Key("port").String()
	ip := cfg.Section("mysql").Key("ip").String()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, ip, port, database)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database " + err.Error())
	}

}
