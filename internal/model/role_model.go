package model

type RoleCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type RoleUpdateRequest struct {
	RoleID      int     `json:"role_id" binding:"required"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type RoleResponse struct {
	RoleID      int    `json:"role_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleListResponse struct {
	Total int64          `json:"total"`
	Items []RoleResponse `json:"items"`
}
