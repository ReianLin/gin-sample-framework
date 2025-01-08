package entity

type Permission struct {
	PermissionID int `json:"permission_id" gorm:"column:permission_id;primaryKey"`
	BaseEntity
	Name string `json:"name" gorm:"column:name"`
	Code string `json:"code" gorm:"column:code"`
	Type int    `json:"type" gorm:"column:type"`
}

func (Permission) TableName() string {
	return "system_permission_list"
}
