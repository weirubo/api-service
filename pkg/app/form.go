package app

import (
	"github.com/gin-gonic/gin"
)

// 入参校验封装

func BindAndValid(c *gin.Context, v interface{}) (bool, error) {
	err := c.ShouldBind(v)
	if err != nil {
		return true, err
	}
	return false, nil
}
