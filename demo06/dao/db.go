package dao

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"lin.com/study/utils"
)

type MyModel struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt utils.LocalDateTimeFormat `gorm:"autoCreateTime;type:datetime"`
	UpdatedAt utils.LocalDateTimeFormat `gorm:"autoUpdateTime;type:timestamp"`
}

func DBGorm() {
	dsn := "root:024680@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile),
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // 日志级别
				Colorful:      true,        // 彩色打印
			},
		),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移数据库表
	// err = db.AutoMigrate(&MyModel{})
	// if err != nil {
	// 	panic("failed to migrate database")
	// }

	// 创建记录
	model := MyModel{Name: "Test Record", ID: 1}
	// db.Create(&model)

	// 输出创建时间和更新时间
	time.Sleep(time.Second)
	// 更新记录
	model.UpdatedAt = utils.LocalDateTimeFormat(time.Now())
	fmt.Printf("Updated At before Update: %+v\n", model.UpdatedAt.String())

	db.Model(&model).Update("Name", "Updated Name")
	// 输出更新时间
	fmt.Printf("Updated At After Update: %+v\n", model)

	resB, err := json.Marshal(&model)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resB))
}
