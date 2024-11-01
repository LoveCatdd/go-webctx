package auth

type JWTConfig struct {
	Jwt struct {
		Secret string `mapstructure:"cecret"`
	} `mapstructure:"jwt"`
}

var JwtConfig JWTConfig
