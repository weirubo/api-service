package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// 读取配置文件
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")   // 设置配置文件的名称
	vp.AddConfigPath("configs/") // 设置配置文件的路径
	vp.SetConfigType("yaml")     // 设置配置文件的类型
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
