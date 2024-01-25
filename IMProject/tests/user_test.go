package tests

import (
	"fmt"
	"testing"

	"lin.com/im/models"
	"lin.com/im/utils"
)

// 测试
func TestUser(t *testing.T) {
	if err := models.DB.AutoMigrate(&models.UserBase{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
	models.DB.Create(&models.UserBase{
		Username: "llll",
		Password: "test",
	})
}
func TestDelete(t *testing.T) {
	res := models.DB.Where("id = ?", 1).Delete(&models.UserBase{})
	if res.RowsAffected == 0 {
		fmt.Println("删除失败")
	} else {
		fmt.Println("删除成功")

	}
}

func TestUpdate(t *testing.T) {
	user := models.UserBase{
		Username: "霖霖",
		Password: "test",
		GORM_MODEL: utils.GORM_MODEL{
			ID: 5,
		},
	}
	res := models.DB.Updates(user)
	if res.RowsAffected == 0 {
		fmt.Println("更新失败")
	} else {
		fmt.Println("更新成功")
	}
}
