package logging

import (
	"fmt"
	"testing"
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

func TestRedConsole(t *testing.T) {
	fmt.Println("\033[31mHello World \033[0m")

	formatter := NewFormatter()

	fmt.Println(formatter.FormatColor("Hello world", ColorGreen))
	fmt.Println(formatter.FormatColor("Hello world", ColorYellow))
	fmt.Println(formatter.FormatColor("Hello world", ColorBlue))
	fmt.Println(formatter.FormatColor("Hello world", ColorMagenta))
	fmt.Println(formatter.FormatColor("Hello world", ColorCyan))
	fmt.Println(formatter.FormatColor("Hello world", ColorWhite))

	
	fmt.Println(formatter.FormatColor("Hello world", ColorBrightRed))
	fmt.Println(formatter.FormatColor("Hello world", ColorBrightGreen))
	fmt.Println(formatter.FormatColor("Hello world", ColorBrightYellow))
	fmt.Println(formatter.FormatColor("Hello world", ColorBrightBlue))
	fmt.Println(formatter.FormatColor("Hello world", ColorBrightMagenta))
	fmt.Println(formatter.FormatColor("Hello world", ColorBrightCyan))
	fmt.Println(formatter.FormatColor("Hello world", ColorBrightWhite))
}
