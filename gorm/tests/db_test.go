package tests

import (
	"log"
	"testing"

	"lin.com/gorm/models"
	"lin.com/gorm/utils"
)

func TestDB(t *testing.T) {
	if err := models.DB.AutoMigrate(&models.User{}); err != nil {
		log.Panicln(err)
	}
	user := models.User{
		Name:     "lin",
		Email:    nil,
		Age:      18,
		Birthday: nil,
	}

	if res := models.DB.Create(&user); res.RowsAffected == 0 {
		log.Println("insert failed")
	} else {
		log.Println("insert success")
		log.Println(user)
	}

}

func TestUpdate(t *testing.T) {
	user := models.User{
		Name: "霖霖",
		Age:  20,
		GORM_MODEL: utils.GORM_MODEL{
			ID: 1,
		},
	}

	if res := models.DB.Updates(user); res.RowsAffected == 0 {
		log.Println("update failed")
	} else {
		log.Println("update success")
		log.Println(user)
	}
}
