package entity

type UserRoleMiddle struct {
	BaseEntity
	UserRoleMiddleId int `json:"user_role_middle_id" gorm:"column:user_role_middle_id;primaryKey"`
	UserID           int `json:"user_id" gorm:"column:user_id"`
	RoleID           int `json:"role_id" gorm:"column:role_id"`
}

func (UserRoleMiddle) TableName() string {
	return "user_role_middle_list"
}
