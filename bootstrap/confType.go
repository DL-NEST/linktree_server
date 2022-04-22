package bootstrap


type Service struct {
	Port       int    `yaml:"port"`
	ServerLog  string `yaml:"serverLog"`
	RequestLog string `yaml:"requestLog"`
	Mode       string `yaml:"mode"`
}

type DB struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	LoginName string `yaml:"loginName"`
	Password  string `yaml:"password"`
}

// Config yaml:""转换工具 https://oktools.net/yaml:""2go
type Config struct {
	Service `yaml:"service"`
	DB      `yaml:"db"`
	Redis struct {
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		LoginName string `yaml:"loginName"`
		Password  string `yaml:"password"`
	} `yaml:"redis"`
	Emqx struct {
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		LoginName string `yaml:"loginName"`
		Password  string `yaml:"password"`
		ClientID  string `yaml:"clientId"`
	} `yaml:"emqx"`
}