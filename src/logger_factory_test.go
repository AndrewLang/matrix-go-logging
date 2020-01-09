package logging

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func defaultLogginConfig() *LoggerConfiguration {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.DebugStyle = LogLevelStyle{"245", "24", ""}
	configuration.InfoStyle = LogLevelStyle{"85", "", "1"}
	configuration.WarnStyle = LogLevelStyle{"226", "124", "4"}
	configuration.ErrorStyle = LogLevelStyle{"166", "232", "1,4"}
	configuration.FatalStyle = LogLevelStyle{"196", "11", "7"}

	return configuration
}
func TestNewLoggerFactory(t *testing.T) {
	factory := NewLoggerFactory()
	length := len(factory.creators)

	assert.Equal(t, 3, length, "There should be 3 creators")
}

func TestConfigureLoggerFactory(t *testing.T) {
	configuration := defaultLogginConfig()

	factory := NewLoggerFactory().Configure(configuration)

	assert.Equal(t, "", factory.Configuration.FileName, "")
	assert.Equal(t, 0, factory.Configuration.MinLevel, "")
	assert.Equal(t, int64(2097152), factory.Configuration.FileSize, "")
	assert.Equal(t, true, factory.Configuration.UseColor, "")
	assert.Equal(t, "Time", factory.Configuration.LayoutNames[0], "")
	assert.Equal(t, "Level", factory.Configuration.LayoutNames[1], "")
	assert.Equal(t, "Name", factory.Configuration.LayoutNames[2], "")
	assert.Equal(t, "Indent", factory.Configuration.LayoutNames[3], "")
	assert.Equal(t, "Message", factory.Configuration.LayoutNames[4], "")
}

func TestCreateLogger(t *testing.T) {
	configuration := defaultLogginConfig()
	factory := NewLoggerFactory().Configure(configuration)
	logger, err := factory.Create(ConsoleLoggerName)

	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, logger, "Logger should not be nil")

	content := "Logger created from factory"
	logger.Debug("Logger configuration: ", logger.GetConfiguration()).
		Debug(content).
		Info(content).
		Warn(content).
		Error(content).
		Fatal(content)
}

func TestCreateLoggerNotFound(t *testing.T) {
	configuration := defaultLogginConfig()
	factory := NewLoggerFactory().Configure(configuration)
	_, err := factory.Create("Mock")

	assert.NotNil(t, err, "Error should not be nil")
}

func TestCreateLoggerWithoutConfig(t *testing.T) {
	factory := NewLoggerFactory()
	_, err := factory.Create(ConsoleLoggerName)

	assert.NotNil(t, err, "Error should not be nil")
}

func TestRegisterCreator(t *testing.T) {

	factory := NewLoggerFactory()
	factory.RegisterCreator("Mock", func(name string) ILogger {
		return NewConsoleLogger(name)
	})
	length := len(factory.creators)
	assert.Equal(t, 4, length, "There should be  creators")
}
