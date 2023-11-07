package main

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Email          string     `validate:"required,email"`
	Age            uint8      `validate:"required,gt=0,lte=100"`
	Gender         string     `validate:"oneof=mal female prefer_not_to"`
	FavouriteColor string     `validate:"iscolor"`
	Addresses      []*Address `validate:"required,dive,required"` // dive对切片和集合元素进行校验
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var validate *validator.Validate

func main() {
	validate = validator.New(validator.WithRequiredStructEnabled())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	validateStruct()
	validateVariable()
}

func validateStruct() {
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "",
		LastName:       "Smith",
		Age:            135,
		Gender:         "male",
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	if err := validate.Struct(user); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Print(err)
			return
		}
		// 判断err是不是validator.ValidationErrors错误
		for _, err := range err.(validator.ValidationErrors) {
			log.Println(err.Namespace())
			log.Println(err.Field())
			log.Println(err.StructNamespace())
			log.Println(err.StructField())
			log.Println(err.Tag())
			log.Println(err.ActualTag())
			log.Println(err.Kind())
			log.Println(err.Type())
			log.Println(err.Value())
			log.Println(err.Param())
		}
		return
	}
}

func validateVariable() {
	myEmail := "joeybloggs.gmail.com"

	if errs := validate.Var(myEmail, "required,email"); errs != nil {
		for err := range errs.(validator.ValidationErrors) {
			log.Println(err)
			return
		}

	}
}
