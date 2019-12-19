package logging

const (
	// OneMega 1M
	OneMega int64 = 1024 *1024 
	// DefaultLogFileSize defautl log file size
	DefaultLogFileSize int64 = OneMega * 2
)

// LoggerConfiguration configuration for logger
type LoggerConfiguration struct {
	LayoutNames []string
	FileName    string
	FileSize    int64
	MinLevel    int
}


// NewLoggerConfiguration create new configuration
func NewLoggerConfiguration(layouts []string) LoggerConfiguration {
	return LoggerConfiguration{
		LayoutNames: layouts,
		FileName:    "",
		FileSize:    DefaultLogFileSize,
		MinLevel:    LevelAll.Value,
	}
}

// ILogger interface for logger
type ILogger interface {
	Configure(config LoggerConfiguration) *ILogger

	StartGroup(name string) *ILogger
	EndGroup() *ILogger
	ResetGroup() *ILogger

	IsEnable(level LogLevel) bool

	Debug(message string, objects ...interface{}) *ILogger
	Info(message string, objects ...interface{}) *ILogger
	Warn(message string, objects ...interface{}) *ILogger
	Error(message string, objects ...interface{}) *ILogger
	Fatal(message string, objects ...interface{}) *ILogger
}
