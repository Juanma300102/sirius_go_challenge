package users

import (
	"time"

	"gorm.io/gorm"
)

type detailUriParameters struct {
	Id int `uri:"id" binding:"required"`
}

type createUserDto struct {
	ID         uint
	Email      string `binding:"required"`
	Password   string `binding:"required"`
	Firsttname string 
	Lastname   string 
	Dni        string 
	Adress     string 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type updateUserDto struct {
	Email      string
	Password   string
	Firstname string 
	Lastname   string 
	Dni        string 
	Adress     string 
}

type queryResult struct {
	Result any
	Err error
}