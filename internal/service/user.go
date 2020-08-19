package service

// 入参校验

type CreateUserRequest struct {
	Name      string `form:"name" json:"name" binding:"max=100"`
	State     uint8  `form:"state,default=1" json:"state" binding:"oneof=0 1"`
	PassWord  string `form:"pass_word" json:"pass_word" binding:"required,min=6,max=10"`
	CreatedBy string `form:"created_by" json:"created_by"`
}

func (svc *Service) CreateUser(param *CreateUserRequest) error {
	return svc.dao.CreateUser(param.Name, param.State, param.PassWord, param.CreatedBy)
}
