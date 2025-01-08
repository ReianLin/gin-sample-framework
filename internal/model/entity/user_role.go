package entity

type UserRole struct {
	BaseEntity
	ID     uint `json:"id" gorm:"column:id;primarykey"`
	UserID int  `json:"user_id" gorm:"column:user_id;type:int;not null;index"`
	RoleID int  `json:"role_id" gorm:"column:role_id;type:int;not null;index"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
