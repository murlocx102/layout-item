package tools

import (
	"go.uber.org/zap/zapcore"
)

const (
	LogTagFieldName = `tag`
)

type Conf struct {
	Debug             bool          //是否启用调试
	Level             zapcore.Level //日志级别
	Encoding          string        //编码格式，目前支持json/console，默认console
	WriteToConsole    bool          //输出到控制台
	WriteToLogFile    bool          //输出到日志文件
	DisableColor      bool          //禁止显示日志颜色
	DisableCaller     bool          //禁止显示调用路径
	EnableShortCaller bool          //显示短调用路径
	EnableStackTrace  bool          //显示堆栈
	CallerSkip        int           //调用堆栈跳过行数
	LogFilePath       string        //日志文件名称，默认为 ./log/app.log，需要事先创建log目录
	LogFileMaxSize    int           //日志文件切割最大大小(单位M),默认100M
	LogFileMaxAge     int           //日志文件保存最大天数，超过此天数的旧日志文件将被删除，默认无限制
	LogFileMaxBackups int           //日志文件最大备份数，超过此备份数的旧日志文件将被删除，默认无限制
	CompressLogFile   bool          //是否压缩日志文件
}

func (c Conf) Normalize() Conf {
	if c.Encoding == "" {
		c.Encoding = `console`
	}

	if c.LogFilePath == "" {
		c.LogFilePath = `./log/app.log`
	}

	return c
}

func (c Conf) DefaultConf() Conf {
	c.Debug = true              // 是否启用调试
	c.Level = -1                // 日志级别,见zapcore.Level, debug:-1
	c.Encoding = "console"      // 编码格式，目前支持json/console，默认console
	c.WriteToConsole = true     // 输出到控制台
	c.WriteToLogFile = false    // 输出到日志文件
	c.DisableColor = false      // 禁止显示日志颜色
	c.DisableCaller = false     // 禁止显示调用路径
	c.EnableShortCaller = false // 显示短调用路径
	c.EnableStackTrace = false  // 显示堆栈
	c.CallerSkip = 0            // 调用堆栈跳过行数
	c.LogFilePath = ""          // 日志文件名称，默认为 ./log/app.log
	c.LogFileMaxSize = 0        // 日志文件切割最大大小(单位M),默认100M
	c.LogFileMaxAge = 0         // 日志文件保存最大天数，超过此天数的旧日志文件将被删除，默认无限制
	c.LogFileMaxBackups = 0     // 日志文件最大备份数，超过此数的就日志文件将被删除，默认无限制
	c.CompressLogFile = false   // 是否压缩日志文件
	return c
}
