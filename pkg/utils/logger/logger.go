package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type debugLogLevel int

const (
	Verbose debugLogLevel = iota + 1
	MoreVerbose
)

var DebugLogLevel int

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

	atom.SetLevel(zap.DebugLevel)

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

func Debugf(level debugLogLevel, format string, args ...interface{}) {
	if int(level) > DebugLogLevel {
		return
	}
	GetLogger().Debugf(format, args...)
}

func Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func Debug(level debugLogLevel, args ...interface{}) {
	if int(level) > DebugLogLevel {
		return
	}
	GetLogger().Debug(args...)
}

func Debugln(level debugLogLevel, args ...interface{}) {
	if int(level) > DebugLogLevel {
		return
	}
	GetLogger().Debugln(args...)
}

func Println(args ...interface{}) {
	fmt.Println(args...)
}
