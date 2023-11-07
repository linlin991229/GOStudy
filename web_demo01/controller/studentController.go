package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	// 校验
	validate := validator.New()
	if err := validate.Struct(student); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var errs string
		for _, err := range err.(validator.ValidationErrors) {
			log.Println(err)
			errs += err.Error() + " "
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}
	if models.SaveStudent(student) {
		c.JSON(200, gin.H{
			"status": "save success",
			"data":   student,
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
