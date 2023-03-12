package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, config := range configs {
		vp.AddConfigPath(config)
	}
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

// 读取配置文件 实现热更新
func (s *Setting) WatchSettingChange() {
	go func() {
		//配置文件监听
		s.vp.WatchConfig()
		//调用重载方法ReloadALLSection来处理热更新的文件监听事件回掉
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadALLSection()
		})
	}()
}
