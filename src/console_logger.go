package logging

import (
	"fmt"
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

	logger.printMessage(content)

	return logger
}

// Info write at info level
func (logger *ConsoleLogger) Info(message string, objects ...interface{}) *ConsoleLogger {

	content := logger.WriteMessage(LevelInfo.Name, message, objects...)

	logger.printMessage(content)

	return logger
}

// Warn write at warn level
func (logger *ConsoleLogger) Warn(message string, objects ...interface{}) *ConsoleLogger {

	content := logger.WriteMessage(LevelWarn.Name, message, objects...)

	logger.printMessage(content)

	return logger
}

// Error write at error level
func (logger *ConsoleLogger) Error(message string, objects ...interface{}) *ConsoleLogger {

	content := logger.WriteMessage(LevelError.Name, message, objects...)

	logger.printMessage(content)

	return logger
}

// Fatal write at fatal level
func (logger *ConsoleLogger) Fatal(message string, objects ...interface{}) *ConsoleLogger {

	content := logger.WriteMessage(LevelFatal.Name, message, objects...)

	logger.printMessage(content)

	return logger
}

func (logger *ConsoleLogger) printMessage(message string) {
	fmt.Println(message)
}
