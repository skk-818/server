package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"server/internal/core/config"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	zap *zap.Logger
}

func NewZapLogger(cfg *config.Logger) (*ZapLogger, error) {
	if cfg == nil {
		return nil, fmt.Errorf("zap config is nil")
	}
	if _, err := os.Stat(cfg.Director); os.IsNotExist(err) {
		if err := os.Mkdir(cfg.Director, os.ModePerm); err != nil {
			return nil, err
		}
	}

	baseLevel := parseLevel(cfg.Level)
	cores := make([]zapcore.Core, 0)

	// 文件输出
	jsonEncoder := getEncoder(cfg, false)
	for _, level := range []zapcore.Level{zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel, zapcore.ErrorLevel} {
		if baseLevel <= level {
			ws := getFileWriter(cfg.Director, level.String())
			core := zapcore.NewCore(jsonEncoder, ws, zap.NewAtomicLevelAt(level))
			cores = append(cores, core)
		}
	}

	// 控制台输出（彩色）
	if cfg.LogInConsole {
		consoleEncoder := getEncoder(cfg, true)
		consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zap.NewAtomicLevelAt(baseLevel))
		cores = append(cores, consoleCore)
	}

	logger := zap.New(zapcore.NewTee(cores...))
	if cfg.ShowLine {
		logger = logger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))
	}

	return &ZapLogger{zap: logger}, nil
}

func (l *ZapLogger) Debug(msg string, fields ...zap.Field) {
	l.zap.Debug(msg, fields...)
}

func (l *ZapLogger) Info(msg string, fields ...zap.Field) {
	l.zap.Info(msg, fields...)
}

func (l *ZapLogger) Warn(msg string, fields ...zap.Field) {
	l.zap.Warn(msg, fields...)
}

func (l *ZapLogger) Error(msg string, fields ...zap.Field) {
	l.zap.Error(msg, fields...)
}

func (l *ZapLogger) Fatal(msg string, fields ...zap.Field) {
	l.zap.Fatal(msg, fields...)
}

func (l *ZapLogger) Sync() error {
	return l.zap.Sync()
}

func (l *ZapLogger) Zap() *zap.Logger {
	return l.zap
}

// ---------- 辅助函数 ----------

func parseLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func getEncoder(cfg *config.Logger, forConsole bool) zapcore.Encoder {
	encCfg := zap.NewProductionEncoderConfig()
	encCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encCfg.EncodeCaller = zapcore.ShortCallerEncoder

	if forConsole && strings.ToLower(cfg.Format) == "console" {
		encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encCfg)
	}
	return zapcore.NewJSONEncoder(encCfg)
}

func getFileWriter(dir, level string) zapcore.WriteSyncer {
	file := filepath.Join(dir, fmt.Sprintf("%s.log", level))
	writer, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(writer)
}
