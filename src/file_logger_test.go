package logging

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	MessageCount   int    = 100000
	MessageContent string = "Log messages to file from test environment for performance test"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

func TestFileLoggerDebugResetIndent(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.FileName = "./logging_test.txt"	
	exception := NewException("Argument null exception")
	logger := NewFileLogger("Test.txt").Configure(configuration)

	deleteIfExists(configuration.FileName)

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
	configuration.FileName = "./logging_slow.txt"
	
	exception := NewException("Argument null exception")
	logger := NewFileLogger("SlowMessageGenerator").Configure(configuration)

	deleteIfExists(configuration.FileName)

	logger.StartGroup("Start Indent").StartGroup("Indent 2")

	for i := 0; i < 90; i++ {
		logger.Info(MessageContent, i, i*2, true, false, exception)
	}
	logger.EndGroup()

	time.Sleep(5 * time.Second)
	local := time.Now()
	logger.Info(local.String())
}

// BenchmarAppendMessagesOneByOne take 325.47s for 1000,000
func BenchmarkAppendMessagesOneByOne(b *testing.B) {

	message := MessageContent

	for i := 0; i < MessageCount; i++ {
		file, err := os.OpenFile("1000-append.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("create file failed", err)
			return
		}

		len, err := file.WriteString(message + "\n")
		if err != nil {
			fmt.Println("Write to file error: ", err, len)
		}
		file.Close()
	}
}

// TestWriteMessagesOneByOne  take 4.82s for 1000,000
func BenchmarkWriteMessagesOneByOne(b *testing.B) {
	fileName := "1000-1-1.txt"
	deleteIfExists(fileName)

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("create file failed", err)
		return
	}
	defer file.Close()

	message := MessageContent

	for i := 0; i < MessageCount; i++ {
		len, err := file.WriteString(message + "\n")
		if err != nil {
			fmt.Println("Write to file error: ", err, len)
		}
	}
}

// TestWriteMessagesBatch take 0.48 for 1000,000
func BenchmarkWriteMessagesBatch(b *testing.B) {
	fileName := "1000-batch.txt"
	deleteIfExists(fileName)

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()

	builder := NewStringBuilder()

	message := MessageContent

	for i := 0; i < MessageCount; i++ {
		builder.AppendLine(message)
	}

	len, err := file.WriteString(builder.String())
	if err != nil {
		fmt.Println("Write to file error: ", err, len)
	}
}

func TestGenerateFileName(t *testing.T) {
	fileName := "1000-batch.txt"
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.FileName = fileName

	logger := NewFileLogger("SlowMessageGenerator").Configure(configuration)

	actual := generateFileName(logger.fileName)

	assert.Equal(t, "1000-batch_1.txt", actual, "New log file name is wrong")
}
