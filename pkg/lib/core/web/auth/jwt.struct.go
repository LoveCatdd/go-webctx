package auth

type JWTConfig struct {
	Jwt struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`
}

var JwtConfig JWTConfig
