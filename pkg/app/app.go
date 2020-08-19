package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应处理

type Response struct {
	C *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	return &Response{
		C: c,
	}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.C.JSON(http.StatusOK, data)
}
