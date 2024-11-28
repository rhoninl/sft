package logger

import (
	"fmt"
)

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
