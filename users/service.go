package users

import (
	"challenge/db"
	"challenge/models"
	"errors"

	"gorm.io/gorm"
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

func Create(dto *models.User) (*models.User, error) {
	conn := db.GetConnection()
	result := conn.Create(&dto)
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