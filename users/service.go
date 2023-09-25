package users

import (
	"challenge/db"
	"challenge/models"
	"errors"

	"gorm.io/gorm"
)

func ListAll(ch chan queryResult) {
	var users []models.User
	conn := db.GetConnection()
	result := conn.Find(&users)
	if result.Error != nil {
		ch <- queryResult{nil, result.Error}
    }
	ch <- queryResult{users, nil}
}

func Create(dto *createUserDto) (*createUserDto, error) {
	conn := db.GetConnection()
	result := conn.Model(&models.User{}).Create(&dto)
	if result.Error != nil {
		return dto, result.Error
    }
	return dto, nil
}

func Retrieve(id int, ch chan queryResult){
	var user models.User
	conn := db.GetConnection()
	result := conn.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ch <- queryResult{user, result.Error}
		}
		ch <- queryResult{user, result.Error}
    }
    ch <- queryResult{user, nil}
}


/* func Update(id int, dto *updateUserDto) (*models.User, error) {
	getOneCh := make(chan queryResult)
	conn := db.GetConnection()
	go Retrieve(id, getOneCh)
	result := <- getOneCh
	if result.Err != nil {
		var t = result.Result.(models.User)
		if errors.Is(result.Err, gorm.ErrRecordNotFound) {
			return &t, result.Err
		}
		return &t, result.Err
    }
	resultQ := conn.Model(&result.Result).Clauses(clause.Returning{}).Updates(dto)
	if resultQ.Error != nil {
		var t = result.Result.(models.User)
		return &t, resultQ.Error
    }
	var t = result.Result.(models.User)
	return &t, nil
} */

/* func Delete(id int) (*models.User, error) {
	conn := db.GetConnection()
	user, err := Retrieve(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, err
		}
		return &user, err
    }
	result := conn.Delete(&user)
	if result.Error != nil {
		return &user, result.Error
    }
	return &user, nil
} */