package logger

import (
	"fmt"
)

type Color string

const (
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Blue   Color = "\033[34m"
	Purple Color = "\033[35m"

	Reset Color = "\033[0m"
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
