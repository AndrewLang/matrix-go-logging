package logging

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/sys/windows"
)

func TestConsoleLogger(t *testing.T) {
	exception := NewException("test")

	if exception.Message != "test" {
		t.Errorf("No message of the exception")
	}
}

func TestConsoleLoggerDebug(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	exception := NewException("Argument null exception")
	logger := NewConsoleLogger("Test").Configure(configuration)

	for i := 0; i < 10; i++ {
		logger.Debug("Log messages to console from test environment", i, i*2, true, false, exception)
	}
}

func TestConsoleLoggerDebugWithIndent(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	exception := NewException("Argument null exception")
	logger := NewConsoleLogger("Test").Configure(configuration)

	logger.StartGroup("Start Indent")

	for i := 0; i < 10; i++ {
		logger.Debug("Log messages to console from test environment", i, i*2, true, false, exception)
	}
	logger.EndGroup()

	for i := 0; i < 10; i++ {
		logger.Debug("Log messages to console from test environment", i, i*2, true, false, exception)
	}

	logger.EndGroup()
}

func TestConsoleLoggerDebugResetIndent(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	exception := NewException("Argument null exception")
	logger := NewConsoleLogger("Test").Configure(configuration)

	logger.StartGroup("Start Indent").StartGroup("Indent 2")

	for i := 0; i < 10; i++ {
		logger.Debug("Log messages to console from test environment", i, i*2, true, false, exception)
	}
	logger.ResetGroup()
}

func TestConfigure(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	logger := NewConsoleLogger("Testing").Configure(configuration)

	logger.Debug("Testing configuration", 1, 2, 3).
		Info("Testing configuration", 1, 2, 3).
		Warn("Testing configuration", 1, 2, 3).
		Error("Testing configuration", 1, 2, 3).
		Fatal("Testing configuration", 1, 2, 3)
}

func TestConfigureLayouts(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	logger := NewConsoleLogger("Testing").Configure(configuration)

	logger.Debug("Testing configuration", 1, 2, 3).
		Info("Testing configuration", 1, 2, 3).
		Warn("Testing configuration", 1, 2, 3).
		Error("Testing configuration", 1, 2, 3).
		Fatal("Testing configuration", 1, 2, 3)
}

func TestWithCustomColor(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.ColorDebug = ColorDarkGray
	configuration.ColorInfo = ColorLightGreen
	configuration.ColorWarn = ColorLightYellow
	configuration.ColorError = ColorLightMagenta
	configuration.ColorFatal = ColorLightRed
	logger := NewConsoleLogger("Testing").Configure(configuration)

	logger.Debug("Testing configuration", 1, 2, 3).
		Info("Testing configuration", 1, 2, 3).
		Warn("Testing configuration", 1, 2, 3).
		Error("Testing configuration", 1, 2, 3).
		Fatal("Testing configuration", 1, 2, 3)
}

func TestWithoutColor(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.UseColor = false
	logger := NewConsoleLogger("Testing").Configure(configuration)

	logger.Debug("Testing configuration", 1, 2, 3).
		Info("Testing configuration", 1, 2, 3).
		Warn("Testing configuration", 1, 2, 3).
		Error("Testing configuration", 1, 2, 3).
		Fatal("Testing configuration", 1, 2, 3)
}

func TestConsole256Color(t *testing.T) {
	fmt.Println("\033[38;5;82mHello \033[38;5;198mWorld\033[0m")

	formatter := NewFormatter()

	for i := 0; i <= 256; i++ {
		fmt.Print(formatter.FormatConsoleWith256Color(PaddingRight(fmt.Sprintf("color %v", i), " ", 12), i))
		if i > 0 && (i+1)%10 == 0 {
			fmt.Println("")
		}
	}

	fmt.Println("")
	fmt.Println("")
}

func TestConsole256BgColor(t *testing.T) {	

	formatter := NewFormatter()

	for i := 0; i <= 256; i++ {
		fmt.Print(formatter.FormatConsoleBgWith256Color(PaddingRight(fmt.Sprintf("color %v", i), " ", 12), i))
		if i > 0 && (i+1)%10 == 0 {
			fmt.Println("")
		}
	}

	fmt.Println("")
	fmt.Println("")
}

func TestColorfulConsole(t *testing.T) {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)

	fmt.Println("\033[31mHello World\033[0m")
	fmt.Println("\x1b[31mHello World\x1b[0m")
	fmt.Println("\u001b[31mHello World\u001b[0m")
	fmt.Println("\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	fmt.Println("\033[1;34mTest color\033[0m")

	formatter := NewFormatter()
	content := "Hello colorful world"

	fmt.Println(formatter.FormatConsoleColor(content, ColorGreen))
	fmt.Println(formatter.FormatConsoleColor(content, ColorYellow))
	fmt.Println(formatter.FormatConsoleColor(content, ColorBlue))
	fmt.Println(formatter.FormatConsoleColor(content, ColorMagenta))
	fmt.Println(formatter.FormatConsoleColor(content, ColorCyan))
	fmt.Println(formatter.FormatConsoleColor(content, ColorWhite))

	fmt.Println(formatter.FormatConsoleColor(content, ColorBrightRed))
	fmt.Println(formatter.FormatConsoleColor(content, ColorBrightGreen))
	fmt.Println(formatter.FormatConsoleColor(content, ColorBrightYellow))
	fmt.Println(formatter.FormatConsoleColor(content, ColorBrightBlue))
	fmt.Println(formatter.FormatConsoleColor(content, ColorBrightMagenta))
	fmt.Println(formatter.FormatConsoleColor(content, ColorBrightCyan))
	fmt.Println(formatter.FormatConsoleColor(content, ColorBrightWhite))
}
