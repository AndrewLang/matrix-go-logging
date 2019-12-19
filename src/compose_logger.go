package logging

// ComposeLogger composed logger
type ComposeLogger struct {
	Loggers []ILogger
}
