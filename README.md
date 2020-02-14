# matrix-go-logging

A logging package for golang.

## Features

Here are the main features and these are what I normally need for the development.

1. Provides ILogger interface to avoide use a concrete logger definition
2. Allows configure the message layout by logging configuration
3. Allows configure log message output destination by logging configuration
4. Allows configure log level to output destination
5. Supports load logging configuration from a file (JSON format)
6. Supports log messages to console window
7. Supports colorize level message to console window
8. Supports log messages to plain text file
9. Supports log messages to json file
10. Support Windows, MacOS and Linux
11. Good performance

## Install

Intall the logging package by run `go get github.com/andrewlang/matrix-go-logging`, then you will be able to create loggers to log messages.

```golang
package main

import (
	"github.com/andrewlang/matrix-go-logging"
)

```

## Examples

### Configure logger factory in code

Configure logging to use Console logger with the message layout: [Datetime] [Level] [Logger Name] [Indent] [Message]

The output message you can see in console window

```golang
[2020-02-13 15:28:08.716] [Info]   [StartClientAction]  Start Actionn to response user's request
[2020-02-13 15:28:08.723] [Debug]  [Test Station]  Test station is started with XXX
```

```golang
import (
	logging "github.com/andrewlang/matrix-go-logging"
)

factory := logging.NewLoggerFactory()
config := logging.NewLogTargetConfigurations()

consoleTarget := logging.NewLogTargetConfiguration("Console", logging.ConsoleLoggerName, []string{logging.Time, logging.Level, logging.Name, logging.Indent, logging.Message})
config.AddTarget(consoleTarget)

factory.Configure(config)

```

### Load configuration from file

Logging factory use a json format file to store the configuration for layouts and loggers

```golang
configFile := NewFile("logging.config.json")
factory := logging.NewLoggerFactory().ConfigureFromFile(LoggingConfigFile)

```

### Create logger instance

After logging factory configured, you can use factory create an instance of ILogger.

```golang
// need a name for the logger
logger, _ := factory.Create("Application")
```


### Only allow output for given log level

Here is an example log all level messages to plain text file and only log error level messages to json file.

```golang
config := NewLogTargetConfigurations()

// log all messages to txt file
fileConfig := NewLogTargetConfiguration("File", FileLoggerName, []string{Time, Level, Name, Indent, Message})
fileConfig.Configuration.FileName = "./test/compose_logger.txt"
config.AddTarget(fileConfig)

// log Error level messages to json file
jsonConfig := NewLogTargetConfiguration("JsonFile", JSONFileLoggerName, []string{Message})
jsonConfig.Configuration.FileName = "./test/compose_logger.json"
jsonConfig.Configuration.MinLevel = LevelError.Value
config.AddTarget(jsonConfig)
```

### Configure max file size

By default, the max log file size is 2M, after reach or over this, a new log file will be generated to keep write messages. You can set the size with following code.

```golang
config := NewLogTargetConfigurations()

// log all messages to txt file
fileConfig := NewLogTargetConfiguration("File", FileLoggerName, []string{Time, Level, Name, Indent, Message})
fileConfig.Configuration.FileName = "./test/logger.txt"
fileConfig.Configuration.FileSize = 5 * 1024 * 1024
config.AddTarget(fileConfig)
```


## Build 

You can beuild the package by

```
go build
```

## Test

You can run unit test by

```
go test -v
```

