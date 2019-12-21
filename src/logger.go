package logging

// Object represent an object
type Object interface{}

// LoggerConfiguration configuration for logger
type LoggerConfiguration struct {
	LayoutNames []string
	FileName    string
	FileSize    int64
	MinLevel    int
}

// NewLoggerConfiguration create new configuration
func NewLoggerConfiguration(layouts []string) *LoggerConfiguration {
	return &LoggerConfiguration{
		LayoutNames: layouts,
		FileName:    Empty,
		FileSize:    DefaultLogFileSize,
		MinLevel:    LevelAll.Value,
	}
}

// ILogger interface for logger
type ILogger interface {
	// Configure configure logger with configuration
	Configure(config *LoggerConfiguration) *ILogger

	// StartGroup start an indent
	StartGroup(name string) *ILogger
	// EndGroup end an indent
	EndGroup() *ILogger
	// ResetGroup reset indent to 0
	ResetGroup() *ILogger

	// IsEnable return a value indicate whether given level is enabled
	IsEnable(level LogLevel) bool

	// Debug write message to debug level
	Debug(message string, objects ...interface{}) *ILogger
	// Info write message to info level
	Info(message string, objects ...interface{}) *ILogger
	// Warn write message to warn level
	Warn(message string, objects ...interface{}) *ILogger
	// Error write message to error level
	Error(message string, objects ...interface{}) *ILogger
	// Fatal write message to fatal level
	Fatal(message string, objects ...interface{}) *ILogger
}
