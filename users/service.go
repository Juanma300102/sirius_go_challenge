package users

import (
	"challenge/db"
	"challenge/models"

	"gorm.io/gorm/clause"
)

func ListAll(ch chan queryResult) {
	var users []models.User
	conn := db.GetConnection()
	result := conn.Find(&users)
	ch <- queryResult{users, result.Error}
}

func Create(dto *createUserDto, ch chan queryResult) {
	conn := db.GetConnection()
	result := conn.Model(&models.User{}).Create(&dto)
	ch <- queryResult{dto, result.Error}
}

func Retrieve(id int, ch chan queryResult){
	var user models.User
	conn := db.GetConnection()
	result := conn.First(&user, id)
    ch <- queryResult{user, result.Error}
}


func Update(id int, dto *updateUserDto, ch chan queryResult){
	conn := db.GetConnection()

	getOneCh := make(chan queryResult)
	go Retrieve(id, getOneCh)
	result := <- getOneCh
	var user = result.Result.(models.User)
	if result.Err != nil {
		ch <- queryResult{&user, result.Err}
    }
	resultQ := conn.Model(&user).Clauses(clause.Returning{}).Updates(dto)
	ch <- queryResult{&user, resultQ.Error}
}

func Delete(id int, ch chan queryResult) {
	conn := db.GetConnection()

	getOneCh := make(chan queryResult)
	go Retrieve(id, getOneCh)
	result := <- getOneCh
	var user = result.Result.(models.User)
	if result.Err != nil {
		ch <- queryResult{&user, result.Err}
    }
	
	resultQ := conn.Delete(&user)
	ch <- queryResult{&user, resultQ.Error}
}