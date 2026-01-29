package logger

import (
	"fiber-template/config"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

// Init initializes Zap with standard Production config + Lumberjack for file rotation.
// Uses Zap's built-in encoder config to avoid nil panic in custom encoder.
func Init(cfg *config.LogConfig) error {
	if cfg == nil {
		cfg = &config.LogConfig{
			Dir:        "./storage/logs",
			Filename:   "app.log",
			MaxDays:    7,
			MaxSizeMB:  100,
			Compress:   true,
			MaxBackups: 7,
		}
	}

	if err := os.MkdirAll(cfg.Dir, 0755); err != nil {
		return err
	}

	fullPath := filepath.Join(cfg.Dir, cfg.Filename)

	// Lumberjack: rotate theo size/ngày, giữ N file, nén file cũ
	ws := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fullPath,
		MaxSize:    cfg.MaxSizeMB,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxDays,
		Compress:   cfg.Compress,
	})

	// Cấu hình chuẩn Zap Production (encoder đầy đủ, không nil)
	zapConfig := zap.NewProductionConfig()
	enc := zapcore.NewJSONEncoder(zapConfig.EncoderConfig)
	core := zapcore.NewCore(enc, ws, zapConfig.Level)

	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return nil
}

// Sync flushes any buffered log entries.
func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}

// Info logs an info message with optional fields.
func Info(msg string, fields ...zap.Field) {
	if Log != nil {
		Log.Info(msg, fields...)
	}
}

// Error logs an error message with optional fields.
func Error(msg string, fields ...zap.Field) {
	if Log != nil {
		Log.Error(msg, fields...)
	}
}

// Warn logs a warning message with optional fields.
func Warn(msg string, fields ...zap.Field) {
	if Log != nil {
		Log.Warn(msg, fields...)
	}
}

// Debug logs a debug message with optional fields.
func Debug(msg string, fields ...zap.Field) {
	if Log != nil {
		Log.Debug(msg, fields...)
	}
}
