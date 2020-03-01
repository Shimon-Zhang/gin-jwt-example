package model

import (
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
