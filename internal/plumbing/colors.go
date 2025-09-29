package plumbing

import "fmt"

type Color string

const (
	ColorReset   Color = "\033[0m"
	ColorRed     Color = "\033[31m"
	ColorGreen   Color = "\033[32m"
	ColorYellow  Color = "\033[33m"
	ColorBlue    Color = "\033[34m"
	ColorMagenta Color = "\033[35m"
	ColorCyan    Color = "\033[36m"
	ColorGray    Color = "\033[37m"
	ColorWhite   Color = "\033[97m"
)

func PrintfColor(color Color, format string, a ...any) {
	fmt.Printf(string(color)+format+string(ColorReset), a...)
}

func SprintfColor(color Color, format string, a ...any) string {
	return fmt.Sprintf(string(color)+format+string(ColorReset), a...)
}

func PrintfRed(format string, a ...any) {
	PrintfColor(ColorRed, format, a...)
}
