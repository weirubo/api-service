package setting

import "time"

// 读取区段配置

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
