package configs

type BaseConfig struct {
	MerchantHTTP HTTP  `yaml:"merchantHTTP"`
	UserRPC      HTTP  `yaml:"userRPC"`
	MYSQL        MYSQL `yaml:"mysql"`
	Redis        Redis `yaml:"redis"`
	Token        Token `yaml:"token"`
	Nats         Nats  `yaml:"nats"`
}

// 开放端口
type HTTP struct {
	Addr      string `yaml:"addr"`      //服务地址
	Port      int    `yaml:"port"`      // 服务端口
	PprofPort int    `yaml:"pprofPort"` // pprof端口
}

// MYSQL mysql配置
type MYSQL struct {
	Addr string `yaml:"addr"` // 地址
	User string `yaml:"user"` // 密码
	Pass string `yaml:"pass"` // 账号
	Db   string `yaml:"db"`   // 数据库
	Slog bool   `yaml:"slog"` // 是否记录gorm操作到日志
}

// redis 配置
type Redis struct {
	Addr string `yaml:"addr"` // 地址
	Pass string `yaml:"pass"` // 密码
	Db   int    `yaml:"db"`   // 库
}

// nats 配置
type Nats struct {
	Addr string `yaml:"addr"` // 地址
	User string `yaml:"user"` // 地址
	Pass string `yaml:"pass"` // 密码
}

// app请求接口的token参数
type Token struct {
	LiveTokenSecret    string `yaml:"liveTokenSecret"`    // token加密密钥
	LiveSessionKey     string `yaml:"liveSessionKey"`     // session的key
	LiveSessionSecret  string `yaml:"liveSessionSecret"`  // session密钥
	LiveTokenStringKey string `yaml:"liveTokenStringKey"` // token字符串key
	LiveExpireTime     int    `yaml:"liveExpireTime"`     // 过期时间
}
