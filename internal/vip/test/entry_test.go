package test

import (
	"go.uber.org/zap"
)

var (
	TestLogger *zap.Logger
)

/* func TestMain(m *testing.M) {
	err := configs.LoadConfing("../../configs", "")
	if err != nil {
		panic("加载配置文件失败" + err.Error())
	}
	cfg := configs.GetConfig()

	loggerConf := logger.Conf{}.DefaultConf()
	logger.NewLogger(&loggerConf, "test") // 初始化日志器

	TestLogger = logger.Logger.With(zap.String("模块", "test"))

	db.Init(cfg) // 初始化db

	os.Exit(m.Run())
} */
