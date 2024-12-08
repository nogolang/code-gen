package configs

type AllConfig struct {
	Mode string `json:"mode"`

	//服务配置
	Server struct {
		ServerName string `json:"serverName"`
		HttpPort   int    `json:"httpPort"`
	}

	//日志配置
	Log struct {
		Level string `json:"level"`
	}

	//数据库配置
	Gorm struct {
		Url string `json:"url"`
	}
}

// 判断开发还是生产环境
func (receiver *AllConfig) IsDev() bool {
	if receiver.Mode == "dev" {
		return true
	} else if receiver.Mode == "prod" {
		return false
	} else if receiver.Mode == "" {
		return true
	}
	return true
}
