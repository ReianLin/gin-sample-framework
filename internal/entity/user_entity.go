package entity

type User struct {
	UserID int `json:"user_id" gorm:"column:user_id;primaryKey;"`
	BaseEntity
	Email    string `json:"email" gorm:"column:email;type:varchar(255);not null;uniqueIndex"`
	Name     string `json:"name" gorm:"column:name;type:varchar(255);not null;uniqueIndex"`
	Account  string `json:"account" gorm:"column:account;type:varchar(255);uniqueIndex"`
	Password string `json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "system_user_list"
}
