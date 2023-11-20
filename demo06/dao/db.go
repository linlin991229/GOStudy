package dao

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lin.com/study/utils"
)

type MyModel struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time                 `gorm:"autoCreateTime"`
	UpdatedAt utils.LocalDateTimeFormat `gorm:"autoUpdateTime"`
}

func DBGorm() {
	dsn := "root:024680@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移数据库表
	// err = db.AutoMigrate(&MyModel{})
	// if err != nil {
	// 	panic("failed to migrate database")
	// }

	// // 创建记录
	model := MyModel{Name: "Test Record"}
	// db.Create(&model)

	// 输出创建时间和更新时间
	fmt.Printf("Created At: %v\n", model.CreatedAt)
	fmt.Printf("Updated At: %v\n", model.UpdatedAt)
	time.Sleep(time.Second)
	// 更新记录

	db.Model(&model).Update("Name", "Updated Name")

	// 输出更新时间
	fmt.Printf("Updated At After Update: %v\n", model.UpdatedAt.String())

}
