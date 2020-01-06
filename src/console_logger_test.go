package logging

import (
	"fmt"
	"testing"

	. "github.com/logrusorgru/aurora"
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

func TestColorfulConsole(t *testing.T) {
	fmt.Println("\033[31mHello World\033[0m")
	fmt.Println("\x1b[31mHello World\x1b[0m")
	fmt.Println("\u001b[31mHello World\u001b[0m")
	fmt.Println("\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	fmt.Println("\033[1;34mTest color\033[0m")

	formatter := NewFormatter()
	content := "Hello colorful world"

	fmt.Println(formatter.FormatColor(content, ColorGreen))
	fmt.Println(formatter.FormatColor(content, ColorYellow))
	fmt.Println(formatter.FormatColor(content, ColorBlue))
	fmt.Println(formatter.FormatColor(content, ColorMagenta))
	fmt.Println(formatter.FormatColor(content, ColorCyan))
	fmt.Println(formatter.FormatColor(content, ColorWhite))

	fmt.Println(formatter.FormatColor(content, ColorBrightRed))
	fmt.Println(formatter.FormatColor(content, ColorBrightGreen))
	fmt.Println(formatter.FormatColor(content, ColorBrightYellow))
	fmt.Println(formatter.FormatColor(content, ColorBrightBlue))
	fmt.Println(formatter.FormatColor(content, ColorBrightMagenta))
	fmt.Println(formatter.FormatColor(content, ColorBrightCyan))
	fmt.Println(formatter.FormatColor(content, ColorBrightWhite))

	fmt.Println("Hello,", Magenta("Aurora"))
	fmt.Println(Bold(Cyan("Cya!")))
}
