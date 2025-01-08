package model

import (
	"gin-sample-framework/internal/entity"
	"time"
)

type UserCreateReq struct {
	Account  string `json:"account" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleIDs  []int  `json:"role_ids"`
}

type UserCreateResp struct {
	UserId int `json:user_id`
}

type UserUpdateReq struct {
	UserID   int     `json:"user_id" binding:"required"`
	Account  *string `json:"account"`
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	Password *string `json:"password"`
	RoleIDs  []int   `json:"role_ids" binding:"required"`
}

type UserDetailResp struct {
	UserRoleDetailDTO
}

type UserDetailListResponse struct {
	Total int64                `json:"total"`
	Items []*UserRoleDetailDTO `json:"items"`
}

type UserRoleDetailDTO struct {
	UserID    int           `json:"user_id" gorm:"column:user_id"`
	Account   string        `json:"account" gorm:"column:account"`
	Email     string        `json:"email" gorm:"column:email"`
	Name      string        `json:"name" gorm:"column:name"`
	CreatedAt time.Time     `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time     `json:"updated_at" gorm:"column:updated_at"`
	Roles     []entity.Role `json:"roles" gorm:"many2many:system_user_role_middle;foreignKey:UserID;joinForeignKey:UserID;References:RoleID;joinReferences:RoleID"`
}
