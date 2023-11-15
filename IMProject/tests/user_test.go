package tests

import (
	"testing"

	"lin.com/im/models"
)

// 测试
func TestUser(t *testing.T) {
	// if err := models.DB.AutoMigrate(&models.UserBase{}); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("success")
	// }
	models.DB.Create(&models.UserBase{
		Username: "test",
		Password: "test",
	})
}
