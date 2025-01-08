package entity

type User struct {
	BaseEntity
	UserID   int    `json:"user_id" gorm:"column:user_id;primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"column:username;uniqueIndex"`
	Password string `json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "user_list"
}
