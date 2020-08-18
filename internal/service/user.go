package service

// 入参校验

type CreateUserRequest struct {
	Name     string `form:"name" binding:"max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
	PassWord string `form:"pass_word" binding:"required,min=10,max=20"`
}
