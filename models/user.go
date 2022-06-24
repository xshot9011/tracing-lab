package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID   string `gorm:"type:uuid;default:uuid_generate_v4()`
	Name string `form:"name"`
	Fibo int    `form:"fibo"`
}

func (u *User) BeforeCreate(tx *gorm.DB) {
	u.ID = uuid.NewV4().String()
}
