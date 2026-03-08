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

// Color 为字符串添加 ANSI 颜色代码。
func Color(color int, str string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, str)
}

// colorize 是一个辅助函数，用于格式化字符串并添加颜色。
func colorize(color int, emoji string, format string, a ...any) string {
	str := fmt.Sprintf(format, a...)
	return Color(color, emoji+" "+str)
}

// Error 使用红色和叉号表情符号格式化错误。
// 它通过 err.Error() 获取对用户友好的错误信息，并能安全地处理包含格式化动词的错误字符串。
func Error(err error) string {
	return colorize(31, "❌", "%s", err.Error())
}

// Err 使用红色和叉号表情符号格式化字符串。
func Err(format string, a ...any) string {
	return colorize(31, "❌", format, a...)
}

// Succ 使用绿色和复选标记表情符号格式化字符串。
func Succ(format string, a ...any) string {
	return colorize(32, "✅", format, a...)
}

// Warn 使用黄色和警告表情符号格式化字符串。
func Warn(format string, a ...any) string {
	return colorize(33, "⚠️", format, a...)
}

// Fatal 使用亮红色和骷髅表情符号格式化字符串。
func Fatal(format string, a ...any) string {
	return colorize(91, "💀", format, a...)
}

// PrintError 将格式化的错误消息打印到 stderr。
func PrintError(prefix, format string, a ...any) {
	p := prefixStr(prefix)
	errLogger.Println(p + Err(format, a...))
}

// PrintSucc 将格式化的成功消息打印到 stdout。
func PrintSucc(prefix, format string, a ...any) {
	p := prefixStr(prefix)
	outLogger.Println(p + Succ(format, a...))
}

// PrintWarn 将格式化的警告消息打印到 stderr。
func PrintWarn(prefix, format string, a ...any) {
	p := prefixStr(prefix)
	errLogger.Println(p + Warn(format, a...))
}

// PrintFatal 将格式化的致命消息打印到 stderr 并退出。
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
