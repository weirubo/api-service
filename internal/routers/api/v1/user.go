package v1

import "github.com/gin-gonic/gin"

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

}
