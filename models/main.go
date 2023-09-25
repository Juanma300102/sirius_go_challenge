package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding="required"`
	Firsttname string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Dni        string `json:"dni"`
	Adress    string `json:"adress"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}