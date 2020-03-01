package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
}

func IsTelephoneExist(phone string) bool {
	var user User
	DB.Where("telephone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func GetUser(phone string) (*User, error) {
	var user User
	DB.Where("telephone=?", phone).First(&user)
	if user.ID == 0 {
		return nil, errors.New("user is not exist")
	}
	return &user, nil

}
