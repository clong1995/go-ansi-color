package pcolor

import (
	"fmt"
	"log"
	"os"
)

var (
	outLogger *log.Logger
	errLogger *log.Logger
)

func init() {
	outLogger = log.New(os.Stdout, "", 0)
	errLogger = log.New(os.Stderr, "", 0)
}

// Color applies an ANSI color code to a string.
func Color(color int, str string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, str)
}

// colorize is a helper function to format and colorize a string.
func colorize(color int, emoji string, format string, a ...any) string {
	str := fmt.Sprintf(format, a...)
	return Color(color, emoji+" "+str)
}

// Error formats an error with a red color and a cross emoji.
// It safely handles error strings that may contain formatting verbs.
func Error(err error) string {
	return colorize(31, "❌", "%s", err.Error())
}

// Err formats a string with a red color and a cross emoji.
func Err(format string, a ...any) string {
	return colorize(31, "❌", format, a...)
}

// Succ formats a string with a green color and a checkmark emoji.
func Succ(format string, a ...any) string {
	return colorize(32, "✅", format, a...)
}

// Warn formats a string with a yellow color and a warning emoji.
func Warn(format string, a ...any) string {
	return colorize(33, "⚠️", format, a...)
}

// Fatal formats a string with a bright red color and a skull emoji.
func Fatal(format string, a ...any) string {
	return colorize(91, "💀", format, a...)
}

// PrintError prints a formatted error message to stderr.
func PrintError(prefix, format string, a ...any) {
	p := prefixStr(prefix)
	errLogger.Println(p + Err(format, a...))
}

// PrintSucc prints a formatted success message to stdout.
func PrintSucc(prefix, format string, a ...any) {
	p := prefixStr(prefix)
	outLogger.Println(p + Succ(format, a...))
}

// PrintWarn prints a formatted warning message to stderr.
func PrintWarn(prefix, format string, a ...any) {
	p := prefixStr(prefix)
	errLogger.Println(p + Warn(format, a...))
}

// PrintFatal prints a formatted fatal message to stderr and exits.
func PrintFatal(prefix, format string, a ...any) {
	p := prefixStr(prefix)
	errLogger.Fatalln(p + Fatal(format, a...))
}

func prefixStr(prefix string) string {
	if prefix != "" {
		return fmt.Sprintf("\x1b[1m[%s]\x1b[0m ", prefix)
	}
	return ""
}
