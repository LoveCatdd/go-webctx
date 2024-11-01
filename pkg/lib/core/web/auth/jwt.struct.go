package auth

type JWTConfig struct {
	Jwt struct {
		Secret string `mapstructure:"scecret"`
	} `mapstructure:"jwt"`
}

var JwtConfig JWTConfig
