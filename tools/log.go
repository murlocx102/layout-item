package tools

import (
	"fmt"
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var rootLogger = NewDebugLogger(``, zapcore.DebugLevel)

//必须先初始化rootLogger，否则调用此方法将抛出空指针错误
func NewLogger(cfg *Conf, tags ...string) *zap.Logger {
	if cfg == nil {
		log.Printf("获取日志配置出错[%v]\n", cfg)
	}

	//初始化日志组件，建议在主程序启动前调用此方法
	rootLogger = MustInitRootLoggerFromCfg(*cfg)

	return rootLogger.With(NewTagField(tags...))
}

//用于调试的日志器
func NewDebugLogger(tag string, level zapcore.Level) *zap.Logger {
	c := zap.NewDevelopmentConfig()

	c.Level = zap.NewAtomicLevelAt(level)
	c.DisableStacktrace = true
	c.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	c.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := c.Build()
	if err != nil {
		panic(fmt.Errorf(`create debug logger err[%v]`, err))
	}

	if tag == "" {
		return logger
	}

	return logger.With(NewTagField(tag))
}

func NewTagField(tags ...string) zap.Field {
	return zap.String(LogTagFieldName, strings.Join(tags, `.`))
}

// 根据日志配置初始化日志组件
func MustInitRootLoggerFromCfg(cfg Conf) *zap.Logger {
	cfg = cfg.Normalize()

	level := zap.NewAtomicLevelAt(cfg.Level)

	encoderCfg := zap.NewDevelopmentEncoderConfig()
	if !cfg.DisableColor {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if !cfg.EnableShortCaller {
		encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
	}

	var (
		options []zap.Option
		allCore []zapcore.Core // 如果需要使用日志系统.请添加对应core,	需处理 encoder,writer,core
	)

	// 如需过滤日志标签,请添加encoder的处理
	var encoder zapcore.Encoder
	if cfg.Encoding == `json` {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}

	// 如果writer不支持并发访问,必须使用锁独占writer
	// 文件是不支持并发写入的. 详情见zapcore.Lock(), zap.CombineWriteSyncers()
	var writers []zapcore.WriteSyncer

	if cfg.WriteToConsole {
		writers = append(writers, zapcore.AddSync(os.Stdout))
	}

	if cfg.WriteToLogFile {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.LogFilePath,
			MaxSize:    cfg.LogFileMaxSize,
			MaxBackups: cfg.LogFileMaxBackups,
			MaxAge:     cfg.LogFileMaxAge,
			Compress:   cfg.CompressLogFile,
		})

		writers = append(writers, w)
	}

	if cfg.Debug {
		options = append(options, zap.Development())
	}

	if !cfg.DisableCaller {
		options = append(options, zap.AddCaller(), zap.AddCallerSkip(cfg.CallerSkip))
	}

	if cfg.EnableStackTrace {
		options = append(options, zap.AddStacktrace(level))
	}

	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writers...), level)

	allCore = append(allCore, core)
	core = zapcore.NewTee(allCore...)

	return zap.New(core, options...)
}

//清空缓存的日志，此方法应在主程序退出前调用
func Flush() error {
	if rootLogger != nil {
		return rootLogger.Sync()
	}

	return nil
}

// 提供给标准库log使用
func StdLogger(logger *zap.Logger) *log.Logger {
	logCh := make(chan string, 100)

	go func() {
		for {
			logStr := <-logCh
			logger.Info(logStr[:len(logStr)-1])
		}
	}()

	return log.New(&Writer{
		LogCh: logCh,
	}, "", log.Lmsgprefix)
}

type Writer struct {
	LogCh chan string
}

func (w *Writer) Write(p []byte) (n int, err error) {
	n = len(p)
	w.LogCh <- string(p)

	return
}
