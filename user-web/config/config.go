package config

type UserServerConfig struct {
	Host string `mapstructure:"targetHost"`
	Port int    `mapstructure:"targetPort"`
}

type JwtConfig struct {
	SigningKey string `mapstructure:"signingKey"`
}

type CaptchaConfig struct {
	Type          string `mapstructure:"type"`
	SourceChinese string `mapstructure:"sourceChinese"`
	EnableCaptcha bool   `mapstructure:"enableCaptcha"`
}

type RedisConfig struct {
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	DB             int    `mapstructure:"db"`
	ExpirationTime string `mapstructure:"expirationTime"`
}

type ServerConfig struct {
	ServerName     string           `mapstructure:"name"`
	ServerPort     int              `mapstructure:"serverPort"`
	UserServerInfo UserServerConfig `mapstructure:"user-srv"`
	JWTInfo        JwtConfig        `mapstructure:"jwt"`
	CaptChaInfo    CaptchaConfig    `mapstructure:"captcha"`
	RedisInfo      RedisConfig      `mapstructure:"redis"`
}
