package handler

import (
	"golang-example/model"
	"golang-example/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(ctx *fiber.Ctx) error {
	users, err := service.FindAllUser()

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": 200,
		"data":   users})
}

func GetSingleUser(ctx *fiber.Ctx) error {
	id, errParseId := strconv.Atoi(ctx.Params("id"))
	if errParseId != nil {
		return errParseId
	}

	user, err := service.FindByUserId(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"data": user})

}

func CreateNewUser(ctx *fiber.Ctx) error {
	// user, err := service.CreateUser()
	payload := new(model.User)

	err := ctx.BodyParser(payload)

	if err != nil {
		return err
	}

	user, errorInsert := service.CreateUser(*payload)
	if errorInsert != nil {
		return errorInsert
	}

	return ctx.JSON(fiber.Map{
		"inserted_record": user})
}

func UpdateUserData(ctx *fiber.Ctx) error {
	id, errParseId := strconv.Atoi(ctx.Params("id"))
	if errParseId != nil {
		return errParseId
	}

	payload := new(model.User)

	err := ctx.BodyParser(payload)

	if err != nil {
		return err
	}

	update, errorUpdate := service.UpdateUser(id, *payload)
	if errorUpdate != nil {
		return errorUpdate
	}

	return ctx.JSON(fiber.Map{
		"updated_record": update})
}

func DeleteUserData(ctx *fiber.Ctx) error {
	id, errParseId := strconv.Atoi(ctx.Params("id"))
	if errParseId != nil {
		return errParseId
	}

	_, errorDeleteUser := service.DeleteUser(id)
	if errorDeleteUser != nil {
		return errorDeleteUser
	}
	return ctx.JSON(fiber.Map{
		"message": "Delete user success"})

}
