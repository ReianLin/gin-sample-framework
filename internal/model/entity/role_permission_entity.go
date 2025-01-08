package entity

type RolePermission struct {
	BaseEntity
	RoleID       int `json:"role_id" gorm:"column:role_id;primaryKey"`
	PermissionID int `json:"permission_id" gorm:"column:permission_id;primaryKey"`
}

func (RolePermission) TableName() string {
	return "role_permission_list"
}
