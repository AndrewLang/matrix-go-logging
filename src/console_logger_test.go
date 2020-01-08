package logging

import (
	"fmt"
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
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

func TestWithCustomColorName(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.ColorDebug = LogLevelStyle{"DarkGray", "", ""}
	configuration.ColorInfo = LogLevelStyle{"LightGreen", "", ""}
	configuration.ColorWarn = LogLevelStyle{"LightYellow", "", ""}
	configuration.ColorError = LogLevelStyle{"LightMagenta", "", ""}
	configuration.ColorFatal = LogLevelStyle{"LightRed", "", ""}

	logger := NewConsoleLogger("Testing").Configure(configuration)

	logger.Debug("Testing configuration", 1, 2, 3).
		Info("Testing configuration", 1, 2, 3).
		Warn("Testing configuration", 1, 2, 3).
		Error("Testing configuration", 1, 2, 3).
		Fatal("Testing configuration", 1, 2, 3)
}

func TestWithCustomColorCode(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.ColorDebug = LogLevelStyle{"237", "", ""}
	configuration.ColorInfo = LogLevelStyle{"36", "", ""}
	configuration.ColorWarn = LogLevelStyle{"226", "", ""}
	configuration.ColorError = LogLevelStyle{"165", "", ""}
	configuration.ColorFatal = LogLevelStyle{"197", "", ""}

	logger := NewConsoleLogger("Testing").Configure(configuration)

	logger.Debug("Testing configuration", 1, 2, 3).
		Info("Testing configuration", 1, 2, 3).
		Warn("Testing configuration", 1, 2, 3).
		Error("Testing configuration", 1, 2, 3).
		Fatal("Testing configuration", 1, 2, 3)
}

func TestWithCombineStyles(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.ColorDebug = LogLevelStyle{"245", "24", ""}
	configuration.ColorInfo = LogLevelStyle{"56", "234", "1"}
	configuration.ColorWarn = LogLevelStyle{"226", "124", "4"}
	configuration.ColorError = LogLevelStyle{"166", "232", "1,4"}
	configuration.ColorFatal = LogLevelStyle{"196", "11", "7"}

	logger := NewConsoleLogger("Testing").Configure(configuration)

	content := "Log message with combined styles"
	logger.Debug(content, 1, 2, 3).
		Info(content, 1, 2, 3).
		Warn(content, 1, 2, 3).
		Error(content, 1, 2, 3).
		Fatal(content, 1, 2, 3)
}

func TestMockWithoutColor(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.UseColor = false
	logger := NewConsoleLogger("Testing").Configure(configuration)

	logger.Debug("Testing configuration", 1, 2, 3).
		Info("Testing configuration", 1, 2, 3).
		Warn("Testing configuration", 1, 2, 3).
		Error("Testing configuration", 1, 2, 3).
		Fatal("Testing configuration", 1, 2, 3)
}

func TestMockConsole256Color(t *testing.T) {
	fmt.Println("\033[38;5;82mHello \033[38;5;198mWorld\033[0m")

	formatter := NewFormatter()

	for i := 0; i <= 256; i++ {
		color := fmt.Sprintf("%s%d%s", Color256Start, i, Color256End)
		fmt.Print(formatter.FormatConsoleOutput(PaddingRight(fmt.Sprintf("Color %v", i), " ", 12), color))

		if i > 0 && (i+1)%10 == 0 {
			fmt.Println("")
		}
	}

	fmt.Println("")
	fmt.Println("")
}

func TestMockConsole256BgColor(t *testing.T) {

	formatter := NewFormatter()

	for i := 0; i <= 256; i++ {
		color := fmt.Sprintf("%s%d%s", Color256BgStart, i, Color256End)
		fmt.Print(formatter.FormatConsoleOutput(PaddingRight(fmt.Sprintf("Color %v", i), " ", 12), color))
		if i > 0 && (i+1)%10 == 0 {
			fmt.Println("")
		}
	}

	fmt.Println("")
	fmt.Println("")
}

func TestMockStyles(t *testing.T) {
	fmt.Println("\033[1m Bold text \033[0m")
	fmt.Println("\033[2m Dim text \033[0m")
	fmt.Println("\033[4m Underline text \033[0m")
	fmt.Println("\033[5m Blink text \033[0m")
	fmt.Println("\033[7m Inverted text \033[0m")
	fmt.Println("\033[8m Hidden text \033[0m")
}

func TestMockCombineStyles(t *testing.T) {
	fmt.Println("\033[1;31m Bold text \033[0m")
	fmt.Println("\033[2;32m Dim text \033[0m")
	fmt.Println("\033[4;33m Underline text \033[0m")
	fmt.Println("\033[5;34m Blink text \033[0m")
	fmt.Println("\033[7;35m Inverted text \033[0m")
	fmt.Println("\033[8;36m Hidden text \033[0m")
	fmt.Println("\033[1;4;31m Bold and underline text \033[0m")
	fmt.Println("")

	fmt.Println("Use 16 colors")
	fmt.Println("\033[30;41m Combined styles with 16 colors\033[0m")
	fmt.Println("\033[31;43;1m Combined styles with 16 colors\033[0m")
	fmt.Println("\033[32;45;4m Combined styles with 16 colors\033[0m")
	fmt.Println("\033[33;41;1;4m Combined styles with 16 colors\033[0m")
	fmt.Println("\033[34;40;7m Combined styles with 16 colors\033[0m")
	fmt.Println("")

	fmt.Println("Use 256 colors")
	fmt.Println("\033[38;5;245;48;5;24m Combined styles with 256 colors\033[0m")
	fmt.Println("\033[38;5;46;48;5;234;1m Combined styles with 256 colors\033[0m")
	fmt.Println("\033[38;5;226;48;5;198;4m Combined styles with 256 colors\033[0m")
	fmt.Println("\033[38;5;166;48;5;232;1;4m Combined styles with 256 colors\033[0m")
	fmt.Println("\033[38;5;196;7m Combined styles with 256 colors\033[0m")
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

	fmt.Println(formatter.FormatConsoleOutput(content, ColorGreen.Name))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorYellow.Name))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBlue.Name))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorMagenta.Name))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorCyan.Name))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorWhite.Name))

	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightRed.Name))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightGreen.Name))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightYellow.Name))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightBlue.Name))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightMagenta.Name))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightCyan.Name))
	fmt.Println(formatter.FormatConsoleOutput(content, ColorBrightWhite.Name))
}

func TestParseLevelStyles(t *testing.T) {
	style := LogLevelStyle{ColorDefaultText.Name, ColorRed.Value, "1 , 4 "}

	actual := style.parseLevelStyles()

	assert.Equal(t, 4, len(actual), "Style length should 4")
	assert.Equal(t, "39", actual[0], "Style length should be 4")
	assert.Equal(t, "48;5;31", actual[1], "Style length should be 4")
	assert.Equal(t, "1", actual[2], "Style length should be 1")
	assert.Equal(t, "4", actual[3], "Style length should be 4")
}
