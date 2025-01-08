package entity

type UserRole struct {
	UserRoleID int `json:"user_role_id" gorm:"column:user_role_id;primaryKey"`
	BaseEntity
	UserID int `json:"user_id" gorm:"column:user_id;uniqueIndex:idx_user_id_role_id"`
	RoleID int `json:"role_id" gorm:"column:role_id;uniqueIndex:idx_user_id_role_id"`
}

func (UserRole) TableName() string {
	return "system_user_role_middle"
}
