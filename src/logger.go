package logging

import (
	// "fmt"
	"strings"
)

// Object represent an object
type Object interface{}

// Any an alias for object
type Any interface{}

// LogLevelStyle level style configuration with foreground, background color and styles
type LogLevelStyle struct {
	Foreground string `json:"foreground"`
	Background string `json:"background"`
	Styles     string `json:"styles"`
}

func (style LogLevelStyle) parseForeground() string {
	color := ""

	if IsNullOrEmpty(style.Foreground) {
		return color
	}

	if IsNumber(style.Foreground) {		
		formatter := NewFormatter()
		number, _ := ToNumber(style.Foreground)
		color = formatter.Format256Color(number)
	} else {
		if value, ok := KnownColors[ForegroundPrefix+style.Foreground]; ok {
			color = value
		}
	}
	return color
}

func (style LogLevelStyle) parseBackground() string {
	color := ""

	if IsNullOrEmpty(style.Background) {
		return color
	}

	if IsNumber(style.Background) {
		formatter := NewFormatter()
		number, _ := ToNumber(style.Background)
		color = formatter.FormatBg256Color(number)
	} else {
		if value, ok := KnownColors[BackgroundPrefix+style.Background]; ok {
			color = value
		}
	}

	return color
}

func (style LogLevelStyle) parseStyle() []string {
	styles := make([]string, 0)

	if IsNullOrEmpty(style.Background) {
		return styles
	}

	for _, part := range strings.Split(style.Styles, StyleSeparator) {
		value := strings.Trim(part, Space)
		if IsNullOrEmpty(value) {
			continue
		}

		if s, ok := KnownStyles[value]; ok {
			styles = append(styles, s)
		} else {
			styles = append(styles, value)
		}
	}

	return styles
}

func (style LogLevelStyle) parseLevelStyles() []string {
	styles := make([]string, 0)
	part := style.parseForeground()

	if NotNullOrEmpty(part) {
		styles = append(styles, part)
	}

	part = style.parseBackground()
	if NotNullOrEmpty(part) {
		styles = append(styles, part)
	}

	parts := style.parseStyle()
	if len(parts) > 0 {
		styles = append(styles, parts...)
	}
	// fmt.Println("Level styles: ", styles)

	return styles
}

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
		ColorDebug:  LogLevelStyle{ColorDefaultText, "", ""},
		ColorInfo:   LogLevelStyle{ColorGreen, "", ""},
		ColorWarn:   LogLevelStyle{ColorYellow, "", ""},
		ColorError:  LogLevelStyle{ColorMagenta, "", ""},
		ColorFatal:  LogLevelStyle{ColorRed, "", ""},
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
