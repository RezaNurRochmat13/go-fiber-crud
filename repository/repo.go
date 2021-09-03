package repository

import (
	"errors"
	"golang-example/config"
	"golang-example/model"

	"gorm.io/gorm"
)

func FindAll() ([]model.User, error) {
	var users []model.User

	db, err := config.GormDatabaseConn()

	if err != nil {
		return nil, err
	}

	db.Find(&users)

	return users, nil

}

func FindById(id int) (model.User, error) {
	var user model.User

	db, err := config.GormDatabaseConn()

	if err != nil {
		return model.User{}, err
	}

	errorNotFound := db.First(&user, id).Error
	errors.Is(errorNotFound, gorm.ErrRecordNotFound)

	return user, errorNotFound
}

func Save(user model.User) (model.User, error) {
	db, err := config.GormDatabaseConn()

	if err != nil {
		return model.User{}, err
	}

	db.Save(&user)

	return user, nil
}

func Delete(user model.User) error {
	db, err := config.GormDatabaseConn()

	if err != nil {
		return err
	}

	db.Delete(&user)

	return nil
}
