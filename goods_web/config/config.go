package config

type ServerConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"name"`
	TimeZone string `mapstructure:"timeZone" json:"timeZone"`
}

type Cfg struct {
	ServerInfo ServerConfig `mapstructure:"server" json:"server"`
	JWTInfo    JwtConfig    `mapstructure:"jwt" json:"jwt"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
}

type JwtConfig struct {
	SigningKey string `mapstructure:"signingKey" json:"signingKey"`
}

type ConsulConfig struct {
	Host             string `mapstructure:"host" json:"host"`
	Port             int    `mapstructure:"port" json:"port"`
	Target           string `mapstructure:"target" json:"target"`
	TargetServerName string `mapstructure:"targetServerName" json:"targetServerName"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Scheme    string `mapstructure:"scheme"`
	Namespace string `mapstructure:"namespace"`
	DataId    string `mapstructure:"dataId"`
	Group     string `mapstructure:"group"`
}
