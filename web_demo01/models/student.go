package models

type Student struct {
	Id       int    `json:"id"`
	Username string `json:"username"  validate:"required"`
	Age      int    `json:"age" validate:"required,gte=0,lte=100"`
}

func SaveStudent(student Student) bool {
	res := DB.Create(&student)
	if res.RowsAffected > 0 {
		return true
	}
	return false
}

func GetStudents() []Student {
	var students []Student
	DB.Find(&students)
	return students
}

func GetStudent(id int) Student {
	var student Student
	DB.Where("id = ?", id).First(&student)
	return student
}

func DeleteStudent(id int) bool {
	res := DB.Where("id = ?", id).Delete(&Student{})
	if res.RowsAffected > 0 {
		return true
	}
	return false
}

func UpdateStudent(student Student) bool {
	res := DB.Model(&student).Updates(student)
	if res.RowsAffected > 0 {
		return true
	}
	return false
}
