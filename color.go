package pcolor

import "fmt"

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
