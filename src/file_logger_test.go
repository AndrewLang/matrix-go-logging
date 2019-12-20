package logging

import (
	"fmt"
	"testing"
)

var (
	MessageCount   int    = 100000
	MessageContent string = "Log messages to file from test environment for performance test"
)

func setupTestCase(t *testing.T) func(t *testing.T, message string) {
	testFolder := "./test"
	createFolder(testFolder)

	return func(t *testing.T, message string) {
		deleted := deleteFolder(testFolder)
		fmt.Println("Delete test folder result: ", deleted, message)
	}
}

func TestFileLoggerDebugResetIndent(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.FileName = "./test/logging_test.txt"
	exception := NewException("Argument null exception")
	logger := NewFileLogger("Test.txt").Configure(configuration)

	teardown := setupTestCase(t)
	defer teardown(t, "File Logger")

	logger.StartGroup("Start Indent").StartGroup("Indent 2")

	for i := 0; i < MessageCount; i++ {
		logger.Info(MessageContent, i, i*2, true, false, exception)
	}
	logger.EndGroup()

	for i := 0; i < 10; i++ {
		logger.Debug(MessageContent, i, i*2, true, false, exception)
	}
	logger.EndGroup()
}

func TestFileLoggerDebugSlowLogging(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.FileName = "./test/logging_slow.txt"

	exception := NewException("Argument null exception")
	logger := NewFileLogger("SlowMessageGenerator").Configure(configuration)

	teardown := setupTestCase(t)
	defer teardown(t, "File logger with slow messages")

	logger.StartGroup("Start Indent").StartGroup("Indent 2")

	for i := 0; i < 90; i++ {
		logger.Info(MessageContent, i, i*2, true, false, exception)
	}
	logger.EndGroup()

	// time.Sleep(5 * time.Second)
	// local := time.Now()
	// logger.Info(local.String())
}
