package model

type UserCreateReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserCreateResp struct {
	UserId int `json:"user_id"`
}
