package pcolor

import (
	"fmt"
	"log"
	"os"

	"github.com/clong1995/go-config"
)

var prefix string
var shortLogger *log.Logger

func init() {
	prefix = config.Value("PREFIX")
	if prefix != "" {
		prefix = "\033[1m[" + prefix + "]\033[0m "
	}
	shortLogger = log.New(os.Stdout, prefix, 0)
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

func PrintError(err error) {
	shortLogger.Println(Error(err))
}

func PrintErr(str string, a ...any) {
	shortLogger.Println(Err(str, a...))
}

func PrintSucc(str string, a ...any) {
	shortLogger.Println(Succ(str, a...))
}

func PrintWarn(str string, a ...any) {
	shortLogger.Println(Warn(str, a...))
}

func PrintFatal(str string, a ...any) {
	shortLogger.Println(Fatal(str, a...))
	os.Exit(1)
}
