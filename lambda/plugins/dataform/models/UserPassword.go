package form

type UserPassword struct {
	ID             int        `gorm:"column:id;primary_key" json:"id"`
	Password       string     `gorm:"column:password" json:"password"`
}

//  TableName sets the insert table name for this struct type
func (u *UserPassword) TableName() string {
	return "users"
}
func (u *UserPassword) GetFromTypes() map[string]string {

	fields := map[string]string{

		"id": "Text",
		"password": "Password",

	}

	return fields
}
