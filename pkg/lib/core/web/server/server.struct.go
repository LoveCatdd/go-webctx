package server

import "github.com/LoveCatdd/util/pkg/lib/core/config"

type App struct {
	Server struct {
		Port string `mapstructure:"port"`
		Name string `mapstructure:"name"`
	} `mapstructure:"server"`
}

func (s *App) FileType() string {
	return config.VIPER_YAML
}

var AppConf = new(App)
