package main

import (
	"fmt"
	"golang-example/config"
	"golang-example/handler"
	"golang-example/model"

	"github.com/gofiber/fiber/v2"
)

func DatabaseMigration() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error", err.Error())
	}

	db.AutoMigrate(
		&model.Person{},
		&model.User{},
		&model.Employee{},
		&model.CreditCard{})
}

func main() {
	DatabaseMigration()

	app := fiber.New()

	app.Get("/users", handler.GetAllUser)
	app.Get("/users/:id", handler.GetSingleUser)
	app.Post("/users", handler.CreateNewUser)
	app.Put("/users/:id", handler.UpdateUserData)
	app.Delete("/users/:id", handler.DeleteUserData)

	app.Listen(":8081")
}
