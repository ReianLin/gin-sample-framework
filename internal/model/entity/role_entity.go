package entity

type Role struct {
	BaseEntity
	RoleID int    `json:"role_id" gorm:"column:role_id;primaryKey"`
	Name   string `json:"name" gorm:"column:name;type:varchar(255);not null;unique"`
}

func (Role) TableName() string {
	return "role_list"
}
