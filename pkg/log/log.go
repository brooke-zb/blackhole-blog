package log

import (
	"blackhole-blog/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Default *zap.SugaredLogger
var Api *zap.SugaredLogger
var Err *zap.SugaredLogger

func getEncoder(encoder string) zapcore.Encoder {
	if encoder == "json" {
		c := zapcore.EncoderConfig{
			TimeKey:        "t",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		return zapcore.NewJSONEncoder(c)
	}
	if encoder == "console" {
		c := zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		return zapcore.NewConsoleEncoder(c)
	}

	panic("invalid encoder with value: " + encoder)
}

func getWriter(writer string, filePath *string) zapcore.WriteSyncer {
	if writer == "stdout" {
		return os.Stdout
	}
	if writer == "file" {
		if filePath != nil {
			logger := &lumberjack.Logger{
				Filename:   *filePath,
				MaxSize:    100,
				MaxBackups: 6,
				MaxAge:     30,
				Compress:   true,
			}
			return zapcore.AddSync(logger)
		}
		panic("connot get log file path")
	}
	panic("invalid writer with value: " + writer)
}

func getLevel(level string) zapcore.Level {
	switch level {
	case "debug", "DEBUG":
		return zap.DebugLevel
	case "info", "INFO":
		return zap.InfoLevel
	case "warn", "WARN":
		return zap.WarnLevel
	case "error", "ERROR":
		return zap.ErrorLevel
	case "dpanic", "DPANIC":
		return zap.DPanicLevel
	case "panic", "PANIC":
		return zap.PanicLevel
	case "fatal", "FATAL":
		return zap.FatalLevel
	}
	return zap.InfoLevel
}

func getLogger(setting setting.LogConfig) *zap.SugaredLogger {
	core := zapcore.NewCore(getEncoder(setting.Encoder), getWriter(setting.Writer, setting.File), getLevel(setting.Level))
	logger := zap.New(core)
	return logger.Sugar()
}

func Setup() {
	Default = getLogger(setting.Config.Log.Default)
	Api = getLogger(setting.Config.Log.Api)
	Err = getLogger(setting.Config.Log.Error)
}
