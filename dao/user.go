package dao

import "github.com/weirubo/api-service/internal/model"

// 封装数据访问

func (d *Dao) CreateUser(name string, state uint8, passWord string, createdBy string) error {
	user := model.User{
		Name:     name,
		State:    state,
		PassWord: passWord,
		Model:    &model.Model{CreatedBy: createdBy},
	}
	return user.Create(d.engine)
}
