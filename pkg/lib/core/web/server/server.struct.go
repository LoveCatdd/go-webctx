package server

import "github.com/LoveCatdd/util/pkg/lib/core/viper"

type App struct {
	Server struct {
		Port string `mapstructure:"port"`
		Name string `mapstructure:"name"`
	} `mapstructure:"server"`
}

func (s *App) FileType() string {
	return viper.VIPER_YAML
}

var AppConf = new(App)
