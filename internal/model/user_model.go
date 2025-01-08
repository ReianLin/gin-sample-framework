package model

type UserCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleIDs  []int  `json:"role_ids"`
}

type UserUpdateRequest struct {
	UserID   int     `json:"user_id" binding:"required"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	RoleIDs  []int   `json:"role_ids"`
}

type UserResponse struct {
	UserID   int            `json:"user_id"`
	Username string         `json:"username"`
	Roles    []RoleResponse `json:"roles"`
}

type UserListResponse struct {
	Total int64          `json:"total"`
	Items []UserResponse `json:"items"`
}
