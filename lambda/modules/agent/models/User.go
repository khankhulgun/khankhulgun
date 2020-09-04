package models

import (

	"time"
)

type User struct {
	ID        int64     `gorm:"column:id;primary_key" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Status        string       `gorm:"column:status;" json:"status"` //type:ENUM('0','1','2')
	Role        int64       `gorm:"column:role" json:"role"`
	Login      string    `gorm:"column:login;unique_index;not null;unique" json:"login"`
	Email      string    `gorm:"column:email;unique_index;not null;unique" json:"email"`
	RegisterNumber      string    `gorm:"column:register_number;not null;unique" json:"register_number"`
	Avatar    string    `gorm:"column:avatar;type:TEXT" json:"avatar"`
	Bio    string    `gorm:"column:bio;type:TEXT" json:"bio"`
	FirstName    string    `gorm:"column:first_name" json:"first_name"`
	LastName    string    `gorm:"column:last_name" json:"last_name"`
	Birthday    time.Time `gorm:"column:birthday;type:DATE" json:"birthday"`
	Phone       string     `gorm:"column:phone" json:"phone"`
	Gender    string    `gorm:"column:gender;" json:"gender"` //type:ENUM('f','m')
	Password    string    `gorm:"column:password" json:"password"`
	FcmToken    string    `gorm:"column:fcm_token" json:"fcm_token"`

}

//  TableName sets the insert table name for this struct type
func (v *User) TableName() string {
	return "users"
}

