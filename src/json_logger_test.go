package logging

import (
	"testing"
	// "github.com/stretchr/testify/assert"
)

// TestJsonFileLoggerDebug test
func TestJsonFileLoggerDebug(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Message})
	configuration.FileName = "./logging_test.json"
	configuration.FileSize = 1024 * 1024 * 5

	exception := NewException("Argument null exception")
	logger := NewJSONFileLogger("JsonLogger").Configure(configuration)

	deleteIfExists(configuration.FileName)

	logger.StartGroup("Start Indent").StartGroup("Indent 2")

	for i := 0; i < 500; i++ {
		logger.Info(MessageContent, i, i*2, true, false, exception)
	}
	logger.EndGroup().EndGroup()
}

func TestJsonFileLoggerLevel(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Message})
	configuration.FileName = "./logging_warn.json"
	configuration.FileSize = 1024 * 1024 * 5
	configuration.MinLevel = LevelWarn.Value
	exception := NewException("Argument null exception")
	logger := NewJSONFileLogger("JsonLogger").Configure(configuration)

	deleteIfExists(configuration.FileName)

	logger.StartGroup("Start Indent").StartGroup("Indent 2")

	for i := 0; i < 500; i++ {
		logger.Info(MessageContent, i, i*2, true, false, exception)
	}

	for i := 0; i < 50; i++ {
		logger.Warn(MessageContent, i, i*2, true, false, exception)
	}

	for i := 0; i < 50; i++ {
		logger.Error(MessageContent, i, i*2, true, false, exception)
	}

	for i := 0; i < 50; i++ {
		logger.Fatal(MessageContent, i, i*2, true, false, exception)
	}


	logger.EndGroup().EndGroup()
}
