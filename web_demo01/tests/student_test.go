package tests

import (
	"fmt"
	"testing"

	"lin.com/web_demo01/models"
)

func TestStudent(t *testing.T) {
	fmt.Println(models.GetStudents())
}
