package auth

import "github.com/LoveCatdd/util/pkg/lib/core/viper"

// VIPER CONFIG STRUCT
type JWTConfig struct {
	Jwt struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`
}

func (j *JWTConfig) FileType() string {
	return viper.VIPER_YAML
}

var JwtConfig = new(JWTConfig)
