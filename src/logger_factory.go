package logging

// LoggerCreator create logger instance
type LoggerCreator func(name string) ILogger

// LoggerFactory logger factory
type LoggerFactory struct {
	Configuration *LoggerConfiguration
	creators      map[string]LoggerCreator
}

// NewLoggerFactory create new logger factory
func NewLoggerFactory() *LoggerFactory {
	factory := &LoggerFactory{}

	factory.creators = make(map[string]LoggerCreator)

	factory.RegisterCreator(ConsoleLoggerName, func(name string) ILogger {
		return NewConsoleLogger(name)
	}).RegisterCreator(FileLoggerName, func(name string) ILogger {
		return NewFileLogger(name)
	}).RegisterCreator(JSONFileLoggerName, func(name string) ILogger {
		return NewJSONFileLogger(name)
	})

	return factory
}

// RegisterCreator register a creator to factory
func (factory *LoggerFactory) RegisterCreator(name string, creator LoggerCreator) *LoggerFactory {
	factory.creators[name] = creator
	return factory
}

// Configure Configure factory
func (factory *LoggerFactory) Configure(config *LoggerConfiguration) *LoggerFactory {
	factory.Configuration = config
	return factory
}

// Create create a logger instance
func (factory *LoggerFactory) Create(name string) (ILogger, error) {
	if factory.Configuration == nil {
		return nil, NewException("Factory configuration is not configured. ")
	}

	creator, ok := factory.creators[name]
	if !ok {
		return nil, NewException("No creator found for logger.")
	}

	logger := creator(name)
	logger.Configure(factory.Configuration)

	return logger, nil
}
