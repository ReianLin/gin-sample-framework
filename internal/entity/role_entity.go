package entity

type Role struct {
	RoleID int `json:"role_id" gorm:"column:role_id;primaryKey"`
	BaseEntity
	Name        string `json:"name" gorm:"column:name;type:varchar(255);not null;uniqueIndex"`
	Description string `json:"description" gorm:"column:description;type:varchar(255);not null"`
}

func (Role) TableName() string {
	return "system_role_list"
}
