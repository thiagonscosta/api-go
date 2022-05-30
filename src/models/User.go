package models

import (
	// "errors"
	// "html"
	// "log"
	// "strings"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID   		string   `gorm:"default:uuid_generate_v4()" json:"id"`
    Name  		string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     	string    `gorm:"size:100;not null;unique" json:"email"`
	Password  	string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt 	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 	time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error 
	if err != nil {
		return &User{}, err
	}
	return u, nil
}