package configs

type BaseConfig struct {
	HTTP  HTTP  `yaml:"http"`
	MYSQL MYSQL `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

type HTTP struct {
	Port      uint `yaml:"port"`      // 服务端口
	PprofPort uint `yaml:"pprofPort"` // pprof端口
}

// MYSQL mysql配置
type MYSQL struct {
	Addr string `yaml:"addr"` // 地址
	User string `yaml:"user"` // 密码
	Pass string `yaml:"pass"` // 账号
	Db   string `yaml:"db"`   // 数据库
}

type Redis struct {
	Addr string `yaml:"addr"` // 地址
	Pass string `yaml:"pass"` // 密码
	Db   int    `yaml:"db"`   // 库
}
