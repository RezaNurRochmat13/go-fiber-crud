package handler

import (
	"fmt"
	"golang-example/config"
	"golang-example/model"
)

func InsertAnimal() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	person := model.Person{Name: "Reja", Email: "reza@gmail.com"}

	db.Create(&person)

	fmt.Println("Success insert using ORM")

}

func InsertUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	user := model.User{Name: "Sandy", Email: "sandy@yahoo.com", Address: "Houston"}

	db.Create(&user)

	fmt.Println("Insert user successfully")
}

func SelectUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var user model.User

	result, err := db.Model(&model.User{}).Rows()

	if err != nil {
		fmt.Errorf("Cannot scan row", err.Error())
	}

	for result.Next() {
		db.ScanRows(result, &user)
		fmt.Println(user.Name)
		fmt.Println(user.Email)
	}

}

func UpdateUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var user model.User

	db.Model(&user).Where("id = ?", 1).Update("name", "Aulia").
		Update("email", "auliad@gmail.com")

	fmt.Println("Update user sucess")
}

func DeleteUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var user model.User

	db.Where("id = ?", 1).Delete(&user)

	fmt.Println("Delete user sucess")
}
