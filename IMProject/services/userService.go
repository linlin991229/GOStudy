package services

import (
	"fmt"

	"lin.com/im/models"
)

type UserService struct{}

func (userService UserService) CreateUser(user *models.UserBase) (bool, error) {
	res := models.DB.Create(user)
	if res.RowsAffected == 0 {
		return false, res.Error
	}
	return true, nil
}
func (userService UserService) DeleteUserById(userId string) (bool, error) {
	if res := models.DB.Where("id = ?", userId).Delete(&models.UserBase{}); res.Error != nil {
		fmt.Println("error")
		return false, res.Error
	}
	return true, nil
}
func (userService UserService) UpdateUser(user *models.UserBase) (bool, error) {
	res := models.DB.Updates(user)
	if res.RowsAffected == 0 {
		return false, res.Error
	}
	return true, nil
}
func (userService UserService) GetUser(userId string) (models.UserBase, error) {
	var user models.UserBase
	res := models.DB.Where("id = ?", userId).First(&user)
	if res.RowsAffected == 0 {
		return user, res.Error
	}
	return user, nil
}
func (userService UserService) GetUsers() ([]models.UserBase, error) {
	var users []models.UserBase
	res := models.DB.Find(&users)
	if res.RowsAffected == 0 {
		return users, res.Error
	}
	return users, nil
}
