package config

type ServerConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"name"`
	TimeZone string `mapstructure:"timeZone" json:"timeZone"`
}

type Cfg struct {
	ServerInfo  ServerConfig  `mapstructure:"server" json:"server"`
	JWTInfo     JwtConfig     `mapstructure:"jwt" json:"jwt"`
	CaptchaInfo CaptchaConfig `mapstructure:"captcha" json:"captcha"`
	RedisInfo   RedisConfig   `mapstructure:"redis" json:"redis"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul" json:"consul"`
}

type JwtConfig struct {
	SigningKey string `mapstructure:"signingKey" json:"signingKey"`
}

type CaptchaConfig struct {
	Type          string `mapstructure:"type" json:"type"`
	SourceChinese string `mapstructure:"sourceChinese" json:"sourceChinese"`
	EnableCaptcha bool   `mapstructure:"enableCaptcha" json:"enableCaptcha"`
}

type RedisConfig struct {
	Host           string `mapstructure:"host" json:"host"`
	Port           int    `mapstructure:"port" json:"port"`
	DB             int    `mapstructure:"db" json:"db"`
	ExpirationTime string `mapstructure:"expirationTime" json:"expirationTime"`
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
