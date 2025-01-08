package entity

type RolePermission struct {
	RolePermissionID int `json:"role_permission_id" gorm:"column:role_permission_id;primaryKey"`
	BaseEntity
	RoleID       int `json:"role_id" gorm:"column:role_id;uniqueIndex:idx_role_permission"`
	PermissionID int `json:"permission_id" gorm:"column:permission_id;uniqueIndex:idx_role_permission"`
}

func (RolePermission) TableName() string {
	return "system_role_permission_middle"
}
