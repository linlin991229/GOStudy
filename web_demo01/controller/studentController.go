package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"lin.com/web_demo01/models"
)

type StudentController struct{}

func (stu StudentController) GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetStudents())

}

func (stu StudentController) GetStudent(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.Atoi(id); err == nil {
		log.Println(err)
		c.JSON(http.StatusOK, models.GetStudent(id))
	}
}

func (stu StudentController) SaveStudent(c *gin.Context) {
	var student models.Student
	err := c.BindJSON(&student)
	if err != nil {
		log.Println(err)
	}
	if models.SaveStudent(student) {
		c.JSON(200, gin.H{
			"status": "save success",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "save failed",
	})
}

func (stu StudentController) UpdateStudent(c *gin.Context) {
	var student models.Student
	err := c.BindJSON(&student)
	if err != nil {
		log.Println(err)
	}
	models.UpdateStudent(student)
}

func (stu StudentController) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return
	}
	if models.DeleteStudent(idNum) {
		c.JSON(200, gin.H{
			"status": "delete success",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "delete failed",
	})
}
