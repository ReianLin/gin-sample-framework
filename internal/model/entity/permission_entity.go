package entity

type Permission struct {
	BaseEntity
	PermissionID   int    `json:"permission_id" gorm:"column:permission_id"`
	PermissionName string `json:"permission_name" gorm:"column:permission_name"`
	PermissionCode string `json:"permission_code" gorm:"column:permission_code"`
	PermissionType int    `json:"permission_type" gorm:"column:permission_type"`
}

func (Permission) TableName() string {
	return "permission_list"
}
