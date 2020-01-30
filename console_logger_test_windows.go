package logging

import (
	"fmt"

	"golang.org/x/sys/windows"
	"testing"
)

func TestColorfulConsole(t *testing.T) {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)

	formatter := NewFormatter()
	content := "Hello colorful world"

	fmt.Println(formatter.FormatConsoleOutput(content, ColorGreen.Value))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorYellow.Value))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBlue.Value))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorMagenta.Value))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorCyan.Value))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorWhite.Value))

	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightRed.Value))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightGreen.Value))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightYellow.Value))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightBlue.Value))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightMagenta.Value))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightCyan.Value))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightWhite.Value))
}
