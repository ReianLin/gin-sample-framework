package entity

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;column:id"`
	Username  string    `json:"username" gorm:"column:username;uniqueIndex"`
	Password  string    `json:"password" gorm:"column:password"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "user_list"
}
