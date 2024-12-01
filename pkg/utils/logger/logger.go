package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var DebugLog = true

type Color string

const (
	// Existing colors
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Blue   Color = "\033[34m"
	Purple Color = "\033[35m"
	Cyan   Color = "\033[36m"
	White  Color = "\033[37m"
	Reset  Color = "\033[0m"

	// Additional colors
	Black        Color = "\033[30m"
	Gray         Color = "\033[90m"
	BrightRed    Color = "\033[91m"
	BrightGreen  Color = "\033[92m"
	BrightYellow Color = "\033[93m"
	BrightBlue   Color = "\033[94m"
	BrightPurple Color = "\033[95m"
	BrightCyan   Color = "\033[96m"
	BrightWhite  Color = "\033[97m"

	// Background colors
	BgBlack  Color = "\033[40m"
	BgRed    Color = "\033[41m"
	BgGreen  Color = "\033[42m"
	BgYellow Color = "\033[43m"
	BgBlue   Color = "\033[44m"
	BgPurple Color = "\033[45m"
	BgCyan   Color = "\033[46m"
	BgWhite  Color = "\033[47m"
)

func WithColor(color Color, message string) string {
	return fmt.Sprintf("%s%s%s", color, message, Reset)
}

func StatusWithColor(status string) string {
	switch status {
	case "Running":
		return WithColor(Green, status)
	case "Pending":
		return WithColor(Yellow, status)
	case "Failed":
		return WithColor(Red, status)
	case "ImagePullBackOff":
		return WithColor(Red, status)
	case "CrashLoopBackOff":
		return WithColor(Red, status)
	case "Error":
		return WithColor(Red, status)
	default:
	}
	return status
}

var (
	zlog *zap.SugaredLogger
)

// InitLogger initializes the logger with default production configuration
func InitLogger() {
	zlog = NewLogger()
	zlog = zlog.WithOptions(zap.AddCallerSkip(1))
}

func NewLogger() *zap.SugaredLogger {
	atom := zap.NewAtomicLevel()

	if DebugLog {
		atom.SetLevel(zap.DebugLevel)
	} else {
		atom.SetLevel(zap.InfoLevel)
	}

	cfg := zap.NewDevelopmentConfig()                                // Plain text format
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Color output for better readability in terminal
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	cfg.Level = atom
	logger := zap.Must(cfg.Build())
	sugaredLogger := logger.Sugar()
	return sugaredLogger
}

// GetLogger returns the global logger instance
func GetLogger() *zap.SugaredLogger {
	if zlog == nil {
		InitLogger()
	}
	return zlog
}

func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	GetLogger().Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}

func Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

func Debugln(args ...interface{}) {
	GetLogger().Debugln(args...)
}

func Infoln(args ...interface{}) {
	GetLogger().Infoln(args...)
}

func Warnln(args ...interface{}) {
	GetLogger().Warnln(args...)
}

func Errorln(args ...interface{}) {
	GetLogger().Errorln(args...)
}

func Fatalln(args ...interface{}) {
	GetLogger().Fatal(args...)
}

func Println(args ...interface{}) {
	fmt.Println(args...)
}
