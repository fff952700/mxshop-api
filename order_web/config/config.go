package config

type ServerConfig struct {
	Host     string   `mapstructure:"host" toml:"host"`
	Port     int      `mapstructure:"port" toml:"port"`
	Name     string   `mapstructure:"name" toml:"name"`
	Id       string   `mapstructure:"id" toml:"id"`
	Tag      []string `mapstructure:"tag"  toml:"tag"`
	TimeZone string   `mapstructure:"timeZone"  toml:"timeZone"`
}

type Cfg struct {
	ServerInfo ServerConfig `mapstructure:"server"  toml:"server"`
	JWTInfo    JwtConfig    `mapstructure:"jwt"  toml:"jwt"`
	ConsulInfo ConsulConfig `mapstructure:"consul"  toml:"consul"`
}

type JwtConfig struct {
	SigningKey string `mapstructure:"signingKey"  toml:"signingKey"`
}

type ConsulConfig struct {
	Host             string `mapstructure:"host"  toml:"host"`
	Port             int    `mapstructure:"port"  toml:"port"`
	Target           string `mapstructure:"target"  toml:"target"`
	TargetServerName string `mapstructure:"targetServerName"  toml:"targetServerName"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Scheme    string `mapstructure:"scheme"`
	Namespace string `mapstructure:"namespace"`
	DataId    string `mapstructure:"dataId"`
	Group     string `mapstructure:"group"`
}
