package service

import (
	"golang-example/model"
	"golang-example/repository"
)

func FindAllUser() ([]model.User, error) {
	result, err := repository.FindAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func FindByUserId(id int) (model.User, error) {
	result, err := repository.FindById(id)
	if err != nil {
		return model.User{}, err
	}

	return result, nil
}

func CreateUser(user model.User) (model.User, error) {
	inserted, err := repository.Save(user)

	if err != nil {
		return model.User{}, err
	}

	return inserted, nil
}

func UpdateUser(id int, payload model.User) (model.User, error) {
	user, errById := repository.FindById(id)
	if errById != nil {
		return model.User{}, errById
	}

	user.Name = payload.Name
	user.Address = payload.Address
	user.Email = payload.Email

	update, errUpdate := repository.Save(user)
	if errUpdate != nil {
		return model.User{}, errById
	}

	return update, nil
}

func DeleteUser(id int) (interface{}, error) {
	user, errById := repository.FindById(id)
	if errById != nil {
		return nil, errById
	}

	errDelete := repository.Delete(user)
	if errDelete != nil {
		return nil, errDelete
	}

	return nil, nil
}
