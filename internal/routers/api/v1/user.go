package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/weirubo/api-service/internal/service"
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
	param := service.CreateUserRequest{}
	response := app.NewResponse(c)
	// 入参校验和绑定
	valid, err := app.BindAndValid(c, &param)
	if valid {
		log.Println(err)
		return
	}

	svc := service.New(c.Request.Context())
	err = svc.CreateUser(&param)
	if err != nil {
		return
	}

	response.ToResponse(gin.H{})
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
	valid, err := app.BindAndValid(c, &param)
	if valid {
		log.Println(err)
		return
	}
	response.ToResponse(gin.H{
		"name":  param.Name,
		"state": param.State,
	})
}
