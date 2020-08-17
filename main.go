package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weirubo/api-service/global"
	"github.com/weirubo/api-service/internal/routers"
	"github.com/weirubo/api-service/pkg/setting"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {
	log.Printf("ServerSetting:%#v", global.ServerSetting)
	gin.SetMode(global.ServerSetting.RunMode) // gin 的运行模式
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

// 读取配置
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	return nil
}
