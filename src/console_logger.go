package logging

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

// ConsoleLogger log to console
type ConsoleLogger struct {
	Name             string
	Formatter        Formatter
	indentLevel      int
	layoutNames      []string
	layoutRepository LayoutRepository
	configuration    *LoggerConfiguration
}

// NewConsoleLogger create new console logger
func NewConsoleLogger(name string) *ConsoleLogger {
	logger := ConsoleLogger{
		Name:             name,
		Formatter:        Formatter{},
		indentLevel:      0,
		layoutNames:      []string{},
		layoutRepository: NewLayoutRepository(),
	}

	/*
	  set console mode to enable virtual terminal processing,
	  otherwise it may not work on some windows
	*/
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)

	return &logger
}

// Configure configure logger
func (logger *ConsoleLogger) Configure(config *LoggerConfiguration) *ConsoleLogger {
	logger.layoutNames = config.LayoutNames
	logger.configuration = config
	return logger
}

// StartGroup start a group
func (logger *ConsoleLogger) StartGroup(name string) *ConsoleLogger {
	logger.indentLevel++
	return logger
}

// EndGroup end a group
func (logger *ConsoleLogger) EndGroup() *ConsoleLogger {
	logger.indentLevel--
	if logger.indentLevel < 0 {
		logger.indentLevel = 0
	}

	return logger
}

// ResetGroup reset
func (logger *ConsoleLogger) ResetGroup() *ConsoleLogger {
	logger.indentLevel = 0
	return logger
}

// IsEnable is given level enabled
func (logger *ConsoleLogger) IsEnable(level LogLevel) bool {
	return level.Value >= logger.configuration.MinLevel
}

// WriteMessage writer message and data to string
func (logger *ConsoleLogger) WriteMessage(level string, message string, objects ...interface{}) string {
	layout := logger.layoutRepository.BuildLayout(logger.layoutNames...)

	logMessage := NewMessage(logger.Name, level, logger.indentLevel, message, objects...)
	content := layout.String(logMessage)

	return content
}

// Debug write at debug level
func (logger *ConsoleLogger) Debug(message string, objects ...interface{}) *ConsoleLogger {

	content := logger.WriteMessage(LevelDebug.Name, message, objects...)

	logger.printMessage(content, LevelDebug)

	return logger
}

// Info write at info level
func (logger *ConsoleLogger) Info(message string, objects ...interface{}) *ConsoleLogger {

	content := logger.WriteMessage(LevelInfo.Name, message, objects...)

	logger.printMessage(content, LevelInfo)

	return logger
}

// Warn write at warn level
func (logger *ConsoleLogger) Warn(message string, objects ...interface{}) *ConsoleLogger {

	content := logger.WriteMessage(LevelWarn.Name, message, objects...)

	logger.printMessage(content, LevelWarn)

	return logger
}

// Error write at error level
func (logger *ConsoleLogger) Error(message string, objects ...interface{}) *ConsoleLogger {

	content := logger.WriteMessage(LevelError.Name, message, objects...)

	logger.printMessage(content, LevelError)

	return logger
}

// Fatal write at fatal level
func (logger *ConsoleLogger) Fatal(message string, objects ...interface{}) *ConsoleLogger {

	content := logger.WriteMessage(LevelFatal.Name, message, objects...)

	logger.printMessage(content, LevelFatal)

	return logger
}

func (logger *ConsoleLogger) printMessage(message string, level LogLevel) {
	if logger.configuration.UseColor {
		style := logger.getLevelStyle(level)
		formatter := NewFormatter()
		fmt.Println(formatter.FormatConsoleOutput(message, style...))
	} else {
		fmt.Println(message)
	}
}

func (logger *ConsoleLogger) getLevelStyle(level LogLevel) []string {
	styles := make([]string, 0)

	switch level.Name {
	case LevelDebug.Name:
		styles = logger.configuration.ColorDebug.parseLevelStyles()
	case LevelInfo.Name:
		styles = logger.configuration.ColorInfo.parseLevelStyles()
	case LevelWarn.Name:
		styles = logger.configuration.ColorWarn.parseLevelStyles()
	case LevelError.Name:
		styles = logger.configuration.ColorError.parseLevelStyles()
	case LevelFatal.Name:
		styles = logger.configuration.ColorFatal.parseLevelStyles()
	}

	return styles
}
