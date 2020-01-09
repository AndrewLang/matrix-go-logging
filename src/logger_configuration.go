package logging

import (
	"encoding/json"
)

// LoggerConfiguration configuration for logger
type LoggerConfiguration struct {
	LayoutNames []string      `json:"layoutNames`
	FileName    string        `json:"fileName"`
	FileSize    int64         `json:"fileSize"`
	MinLevel    int           `json:"minLevel"`
	UseColor    bool          `json:"useColor"`
	DebugStyle  LogLevelStyle `json:"debugStyle"`
	InfoStyle   LogLevelStyle `json:"infoStyle"`
	WarnStyle   LogLevelStyle `json:"warnStyle"`
	ErrorStyle  LogLevelStyle `json:"errorStyle"`
	FatalStyle  LogLevelStyle `json:"fatalStyle"`
}

// NewLoggerConfiguration create new configuration
func NewLoggerConfiguration(layouts []string) *LoggerConfiguration {
	return &LoggerConfiguration{
		LayoutNames: layouts,
		FileName:    Empty,
		FileSize:    DefaultLogFileSize,
		MinLevel:    LevelAll.Value,
		UseColor:    true,
		DebugStyle:  LogLevelStyle{ColorDefaultText.Name, "", ""},
		InfoStyle:   LogLevelStyle{ColorGreen.Name, "", ""},
		WarnStyle:   LogLevelStyle{ColorYellow.Name, "", ""},
		ErrorStyle:  LogLevelStyle{ColorMagenta.Name, "", ""},
		FatalStyle:  LogLevelStyle{ColorRed.Name, "", ""},
	}
}

// ToJSON Convert configuration to json string
func (config *LoggerConfiguration) ToJSON() string {
	content, _ := json.MarshalIndent(config, "", "\t")

	return string(content)
}

// FromJSON Load configuration from json string
func (config *LoggerConfiguration) FromJSON(jsonContent string) {
	json.Unmarshal([]byte(jsonContent), config)
}

/*=======================================================================*/

// LogTargetConfiguration configure log output
type LogTargetConfiguration struct {
	Name          string               `json:"name"`
	Configuration *LoggerConfiguration `json:"config"`
}

// NewLogTargetConfiguration create a new log target configuration
func NewLogTargetConfiguration(name string, layouts []string) *LogTargetConfiguration {
	return &LogTargetConfiguration{
		Name:          name,
		Configuration: NewLoggerConfiguration(layouts),
	}
}

// ToJSON serialize to json string
func (config *LogTargetConfiguration) ToJSON() string {
	content, _ := json.MarshalIndent(config, "", "\t")

	return string(content)
}

// FromJSON deserialize from json string
func (config *LogTargetConfiguration) FromJSON(jsonContent string) {
	json.Unmarshal([]byte(jsonContent), config)
}
