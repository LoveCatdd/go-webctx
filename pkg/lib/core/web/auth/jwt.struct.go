package auth

// VIPER CONFIG STRUCT
type JWTConfig struct {
	Jwt struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`
}

func (j *JWTConfig) FileType() string {
	return "yaml"
}

var JwtConfig = new(JWTConfig)
