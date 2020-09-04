package models

import (

	"time"
)

type PasswordReset struct {
	Token      string    `gorm:"column:token;not null;" json:"token"`
	Wrong      int    `gorm:"column:wrong;not null;" json:"wrong"`
	Email      string    `gorm:"column:email;primary_key;" json:"email"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
}

//  TableName sets the insert table name for this struct type
func (v *PasswordReset) TableName() string {
	return "password_resets"
}

