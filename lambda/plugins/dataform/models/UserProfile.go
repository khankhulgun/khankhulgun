package form

import "time"

type UserProfile struct {
	Avatar         string     `gorm:"column:avatar" json:"avatar"`
	Bio            string     `gorm:"column:bio" json:"bio"`
	Birthday       time.Time `gorm:"column:birthday" json:"birthday"`
	Phone       string `gorm:"column:phone" json:"phone"`
	FirstName      string     `gorm:"column:first_name" json:"first_name"`
	Gender         string     `gorm:"column:gender" json:"gender"`
	ID             int        `gorm:"column:id;primary_key" json:"id"`
	LastName       string     `gorm:"column:last_name" json:"last_name"`
	Login      string    `gorm:"column:login;unique_index;not null;unique" json:"login"`
	Email      string    `gorm:"column:email;unique_index;not null;unique" json:"email"`
	RegisterNumber      string    `gorm:"column:register_number;not null;unique" json:"register_number"`

}

//  TableName sets the insert table name for this struct type
func (u *UserProfile) TableName() string {
	return "users"
}

func (u *UserProfile) GetFromTypes() map[string]string {

	fields := map[string]string{
		"id": "Text",
		"email": "Text",
		"login": "Text",
		"register_number": "Text",
		"first_name": "Text",
		"last_name": "Text",
		"birthday": "Date",
		"gender": "Radio",
		"avatar": "Text",
		"phone": "Text",

	}

	return fields
}