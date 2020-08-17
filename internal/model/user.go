package model

// 创建 user model
type User struct {
	*Model
	Name     string `json:"name"`
	State    uint8  `json:"state"`
	PassWord string `json:"pass_word"`
}

// 定义方法
func (u User) TableName() string {
	return "t_user"
}
