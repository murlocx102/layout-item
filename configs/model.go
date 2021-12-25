package configs

type Base struct {
	HTTP HTTP `json:"http"`
}

type HTTP struct {
	Port      uint `json:"port"`      // 服务端口
	PprofPort uint `json:"pprofPort"` // pprof端口
}
