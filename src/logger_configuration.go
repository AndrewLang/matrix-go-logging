package logging

// LoggerConfiguration configuration for logger
type LoggerConfiguration struct {
	LayoutNames []string
	FileName    string
	FileSize    int64
	MinLevel    int
	UseColor    bool
	ColorDebug  LogLevelStyle
	ColorInfo   LogLevelStyle
	ColorWarn   LogLevelStyle
	ColorError  LogLevelStyle
	ColorFatal  LogLevelStyle
}

// NewLoggerConfiguration create new configuration
func NewLoggerConfiguration(layouts []string) *LoggerConfiguration {
	return &LoggerConfiguration{
		LayoutNames: layouts,
		FileName:    Empty,
		FileSize:    DefaultLogFileSize,
		MinLevel:    LevelAll.Value,
		UseColor:    true,
		ColorDebug:  LogLevelStyle{ColorDefaultText.Name, "", ""},
		ColorInfo:   LogLevelStyle{ColorGreen.Name, "", ""},
		ColorWarn:   LogLevelStyle{ColorYellow.Name, "", ""},
		ColorError:  LogLevelStyle{ColorMagenta.Name, "", ""},
		ColorFatal:  LogLevelStyle{ColorRed.Name, "", ""},
	}
}
