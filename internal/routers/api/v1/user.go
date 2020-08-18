package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/weirubo/api-service/pkg/app"
)

// 路由的 Handler

type User struct {
}

func NewUser() User {
	return User{}
}

// 定义方法
// 增
func (u User) Create(c *gin.Context) {

}

// 删
func (u User) Delete(c *gin.Context) {

}

// 改
func (u User) Update(c *gin.Context) {

}

// 查
func (u User) Get(c *gin.Context) {

}

// 查询列表
func (u User) List(c *gin.Context) {
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid == true {
		log.Fatalf("valid err:%s", errs)
		return
	}
	response.ToResponse(gin.H{
		"name":  param.Name,
		"state": param.State,
	})
	return
}
