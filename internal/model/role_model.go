package model

import (
	"gin-sample-framework/internal/entity"
	"time"
)

type RolePermissionDetailDTO struct {
	RoleID      int                 `json:"role_id" gorm:"column:role_id"`
	Name        string              `json:"name" gorm:"column:name"`
	Description string              `json:"description" gorm:"column:description"`
	CreatedAt   time.Time           `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time           `json:"updated_at" gorm:"column:updated_at"`
	Permissions []entity.Permission `json:"permissions" gorm:"many2many:system_role_permission_middle;foreignKey:RoleID;joinForeignKey:RoleID;References:PermissionID;joinReferences:PermissionID"`
}

type RoleCreateRequest struct {
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description"`
	PermissionIDs []int  `json:"permission_ids" binding:"required"`
}

type RoleCreateResponse struct {
	RoleID int `json:"role_id"`
}

type RoleUpdateRequest struct {
	RoleID        int     `json:"role_id" binding:"required"`
	Name          *string `json:"name"`
	Description   *string `json:"description"`
	PermissionIDs []int   `json:"permission_ids" binding:"required"`
}

type RoleDetailResponse struct {
	RolePermissionDetailDTO
}

type RoleDetailListResponse struct {
	Total int64                      `json:"total"`
	Items []*RolePermissionDetailDTO `json:"items"`
}
