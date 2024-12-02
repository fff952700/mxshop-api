package config

type UserServerConfig struct {
	Host        string
	Port        int
	ServiceName string
}

type Cfg struct {
	ServerConfig   string `mapstructure:"serverNme" json:"serverNme"`
	ServerPort     int    `mapstructure:"serverPort" json:"serverPort"`
	TimeZone       string `mapstructure:"timeZone" json:"timeZone"`
	UserServerInfo UserServerConfig
	JWTInfo        JwtConfig     `mapstructure:"jwt" json:"jwt"`
	CaptchaInfo    CaptchaConfig `mapstructure:"captcha" json:"captcha"`
	RedisInfo      RedisConfig   `mapstructure:"redis" json:"redis"`
	ConsulInfo     ConsulConfig  `mapstructure:"consul" json:"consul"`
}

type JwtConfig struct {
	SigningKey string `mapstructure:"SigningKey" json:"SigningKey"`
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
