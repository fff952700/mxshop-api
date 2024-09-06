package config

type UserServerConfig struct {
	Host string `mapstructure:"target-host"`
	Port int    `mapstructure:"target-port"`
}

type JwtConfig struct {
	SigningKey string `mapstructure:"SigningKey"`
}

type ServerConfig struct {
	ServerName     string           `mapstructure:"name"`
	ServerPort     int              `mapstructure:"server-port"`
	UserServerInfo UserServerConfig `mapstructure:"user-srv"`
	JWTInfo        JwtConfig        `mapstructure:"jwt"`
}
