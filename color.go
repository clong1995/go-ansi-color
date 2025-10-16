package pcolor

import (
	"fmt"
	"log"
	"os"
)

var shortLogger *log.Logger

func init() {
	shortLogger = log.New(os.Stdout, "", 0)
}

func Color(color int, str string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, str)
}

func Error(err error) string {
	return Color(31, "❌ "+err.Error())
}

func Err(str string, a ...any) string {
	str = fmt.Sprintf(str, a...)
	return Color(31, "❌ "+str)
}

func Succ(str string, a ...any) string {
	str = fmt.Sprintf(str, a...)
	return Color(32, "✅ "+str)
}

func Warn(str string, a ...any) string {
	str = fmt.Sprintf(str, a...)
	return Color(33, "⚠️ "+str)
}

func Fatal(str string, a ...any) string {
	str = fmt.Sprintf(str, a...)
	return Color(91, "💀 "+str)
}

func PrintError(prefix, str string, a ...any) {
	p := prefixStr(prefix)
	shortLogger.Println(p + Err(str, a...))
}

func PrintSucc(prefix, str string, a ...any) {
	p := prefixStr(prefix)
	shortLogger.Println(p + Succ(str, a...))
}

func PrintWarn(prefix, str string, a ...any) {
	p := prefixStr(prefix)
	shortLogger.Println(p + Warn(str, a...))
}

func PrintFatal(prefix, str string, a ...any) {
	p := prefixStr(prefix)
	shortLogger.Fatalln(p + Fatal(str, a...))
}

func prefixStr(prefix string) string {
	if prefix != "" {
		return "\033[1m[" + prefix + "]\033[0m "
	}
	return ""
}
