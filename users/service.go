package users

import (
	"challenge/db"
	"challenge/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ListAll() ([]models.User, error) {
	var users []models.User
	conn := db.GetConnection()
	result := conn.Find(&users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
    }
    return users, nil
}

func Create(dto *createUserDto) (*createUserDto, error) {
	conn := db.GetConnection()
	result := conn.Model(&models.User{}).Create(&dto)
	if result.Error != nil {
		return dto, result.Error
    }
	return dto, nil
}

func Retrieve(id int) (models.User, error) {
	var user models.User
	conn := db.GetConnection()
	result := conn.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, nil
		}
		return user, result.Error
    }
    return user, nil
}


func Update(id int, dto *updateUserDto) (*models.User, error) {
	conn := db.GetConnection()
	user, err := Retrieve(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, nil
		}
		return &user, err
    }
	result := conn.Model(&user).Clauses(clause.Returning{}).Updates(dto)
	if result.Error != nil {
		return &user, result.Error
    }
	return &user, nil
}