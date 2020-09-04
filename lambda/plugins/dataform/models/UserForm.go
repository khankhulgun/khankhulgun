package form

import "time"

type UserForm struct {
	Avatar         string     `gorm:"column:avatar" json:"avatar"`
	Bio            string     `gorm:"column:bio" json:"bio"`
	Birthday       time.Time 	  `gorm:"column:birthday" json:"birthday"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"-"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"-"`
	Email          string     `gorm:"column:email" json:"email"`
	FcmToken       string     `gorm:"column:fcm_token" json:"fcm_token"`
	FirstName      string     `gorm:"column:first_name" json:"first_name"`
	Gender         string     `gorm:"column:gender" json:"gender"`
	ID             int        `gorm:"column:id;primary_key" json:"id"`
	LastName       string     `gorm:"column:last_name" json:"last_name"`
	Login          string     `gorm:"column:login" json:"login"`
	Password       string     `gorm:"column:password" json:"password"`
	Phone       string     `gorm:"column:phone" json:"phone"`
	RegisterNumber string     `gorm:"column:register_number" json:"register_number"`
	Role           int        `gorm:"column:role" json:"role"`
	//Status         string     `gorm:"column:status" json:"status"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"-"`
}

//  TableName sets the insert table name for this struct type
func (u *UserForm) TableName() string {
	return "users"
}

func (u *UserForm) GetFromTypes() map[string]string {

	fields := map[string]string{

		"id": "Text",
		"password": "Password",
		"birthday": "Date",

	}

	return fields
}
