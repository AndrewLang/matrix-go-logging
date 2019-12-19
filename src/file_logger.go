package logging

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// FileLogger log message to file
type FileLogger struct {
	Name             string
	Formatter        Formatter
	indentLevel      int
	layoutNames      []string
	fileName         string
	fileSize         int64
	layoutRepository LayoutRepository
	buffer           *StringBuilder
	channel          chan string
	mutex            sync.Mutex
	file             *os.File
	timer            *time.Timer
	configuration    LoggerConfiguration
}

// NewFileLogger create new file logger
func NewFileLogger(name string) *FileLogger {
	logger := &FileLogger{
		Name:             name,
		Formatter:        Formatter{},
		indentLevel:      0,
		layoutNames:      []string{},
		fileName:         "",
		fileSize:         1024 * 1024 * 2,
		layoutRepository: NewLayoutRepository(),
		buffer:           NewStringBuilder(),
		channel:          make(chan string),
	}
	runtime.SetFinalizer(logger, closeFileLogger)

	logger.timer = time.AfterFunc(3*time.Second, func() {
		fmt.Println("Flush message to file by timer")
		logger.writeFile()
	})

	return logger
}

// Configure configure logger
func (logger *FileLogger) Configure(config LoggerConfiguration) *FileLogger {
	logger.layoutNames = config.LayoutNames
	logger.fileName = config.FileName
	logger.fileSize = config.FileSize
	logger.configuration = config
	return logger
}

// StartGroup start a group
func (logger *FileLogger) StartGroup(name string) *FileLogger {
	logger.indentLevel++
	fmt.Println("Current indent", logger.indentLevel)
	return logger
}

// EndGroup end a group
func (logger *FileLogger) EndGroup() *FileLogger {
	logger.indentLevel--
	if logger.indentLevel < 0 {
		logger.indentLevel = 0
	}
	fmt.Println("Current indent", logger.indentLevel)
	return logger
}

// ResetGroup reset
func (logger *FileLogger) ResetGroup() *FileLogger {
	logger.indentLevel = 0
	fmt.Println("Current indent", logger.indentLevel)
	return logger
}

// IsEnable is given level enabled
func (logger *FileLogger) IsEnable(level LogLevel) bool {
	return level.Value >= logger.configuration.MinLevel
}

// WriteMessage writer message and data to string
func (logger *FileLogger) WriteMessage(level string, message string, objects ...interface{}) string {
	layout := logger.layoutRepository.BuildLayout(logger.layoutNames...)

	logMessage := NewMessage(logger.Name, level, logger.indentLevel, message, objects...)
	content := layout.String(logMessage)

	return content
}

// Debug write at debug level
func (logger *FileLogger) Debug(message string, objects ...interface{}) *FileLogger {

	content := logger.WriteMessage(LevelDebug.Name, message, objects...)

	logger.printMessage(content)

	return logger
}

// Info write at info level
func (logger *FileLogger) Info(message string, objects ...interface{}) *FileLogger {

	content := logger.WriteMessage(LevelInfo.Name, message, objects...)

	logger.printMessage(content)

	return logger
}

// Warn write at warn level
func (logger *FileLogger) Warn(message string, objects ...interface{}) *FileLogger {

	content := logger.WriteMessage(LevelWarn.Name, message, objects...)

	logger.printMessage(content)

	return logger
}

// Error write at error level
func (logger *FileLogger) Error(message string, objects ...interface{}) *FileLogger {

	content := logger.WriteMessage(LevelError.Name, message, objects...)

	logger.printMessage(content)

	return logger
}

// Fatal write at fatal level
func (logger *FileLogger) Fatal(message string, objects ...interface{}) *FileLogger {

	content := logger.WriteMessage(LevelFatal.Name, message, objects...)

	logger.printMessage(content)

	return logger
}

func (logger *FileLogger) writeFile() {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	size, err := getFileSize(logger.fileName)
	if err != nil {
		return
	}
	if size >= logger.fileSize {
		logger.file = nil
		logger.fileName = generateFileName(logger.configuration.FileName)
		logger.initialize()
	}

	content := logger.buffer.String()
	logger.buffer.Reset()

	_, err = logger.file.WriteString(content)
	if err != nil {
		fmt.Println("Write to file error: ", err)
		return
	}
}

func (logger *FileLogger) initialize() {
	if logger.file == nil {
		file, err := os.OpenFile(logger.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err == nil {
			logger.file = file
		} else {
			fmt.Println("Create file writer error: ", err, logger.fileName)
		}
	}

}

func closeFileLogger(logger *FileLogger) {
	logger.buffer.Reset()
	if logger.file != nil {
		logger.file.Sync()
		logger.file.Close()
		logger.file = nil
	}
	logger.timer.Stop()

	fmt.Println("File logger closed")
}

func (logger *FileLogger) printMessage(message string) {
	logger.initialize()

	logger.buffer.AppendLine(message)

	if logger.buffer.Lines > 100 {
		// fmt.Println("Flush to file with count achieve 100")
		logger.writeFile()
	}
}
