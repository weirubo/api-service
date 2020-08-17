package main

import (
	"net/http"
	"time"

	"github.com/weirubo/api-service/internal/routers"
)

func main() {
	router := routers.NewRouter()

	// 自定义 http.Server
	s := &http.Server{
		Addr:         ":8080",          // 监听端口
		Handler:      router,           // 处理程序
		ReadTimeout:  10 * time.Second, // 允许读取的最大时间
		WriteTimeout: 10 * time.Second, // 允许写入的最大时间
		// MaxHeaderBytes: 1 << 20,          // 请求头的最大字节数
		MaxHeaderBytes: 1024 * 1024, // 请求头的最大字节数
	}
	s.ListenAndServe()
}
