package logging

import (
	"fmt"
)

const (
	// LeftBracket left bracket
	LeftBracket = "["
	// RightBracket right bracket
	RightBracket = "]"

	// Name name
	Name = "Name"
	// Time time
	Time = "Time"
	// Level level
	Level = "Level"
	// Indent indent
	Indent = "Indent"
	// Message message
	Message = "Message"
	// Compose compose
	Compose = "Compose"

	// NameLayout name layout
	NameLayout = "NameLayout"
	// TimeLayout time layout
	TimeLayout = "TimeLayout"
	// LevelLayout level
	LevelLayout = "LevelLayout"
	// IndentLayout indent
	IndentLayout = "IndentLayout"
	// MessageLayout message
	MessageLayout = "MessageLayout"
	// ComposeLayout compose
	ComposeLayout = "ComposeLayout"
)

// ======================================================================================

// LogMessage message for logging
type LogMessage struct {
	Name     string
	Level    string
	Message  string
	Datetime string
	Indent   int
	Data     []interface{}
}

// NewMessage create a message
func NewMessage(name string, level string, indent int, message string, data ...interface{}) LogMessage {
	formatter := Formatter{}
	contextData := make([]interface{}, 0)
	for _, item := range data {
		contextData = append(contextData, item)
	}

	return LogMessage{
		Name:     name,
		Level:    level,
		Message:  message,
		Datetime: formatter.FormatNow(),
		Indent:   indent,
		Data:     contextData,
	}
}

// ILogLayout define layout of logging
type ILogLayout interface {
	String(message LogMessage) string
}

// ======================================================================================

// LogTimeLayout get log time
type LogTimeLayout struct {
	date string
}

// NewLogTimeLayout create a LogTimeLayout
func NewLogTimeLayout() *LogTimeLayout {
	formatter := Formatter{}
	return &LogTimeLayout{formatter.FormatNow()}
}

func (l *LogTimeLayout) String(message LogMessage) string {
	return fmt.Sprintf("%s%s%s", LeftBracket, l.date, RightBracket)
}

// ======================================================================================

// LogLevelLayout level layout
type LogLevelLayout struct {
	Padding int
}

// NewLogLevelLayout create level layout
func NewLogLevelLayout() *LogLevelLayout {
	return &LogLevelLayout{8}
}

func (l *LogLevelLayout) String(message LogMessage) string {
	// format use right padding
	result := fmt.Sprintf("%s%s%s", LeftBracket, message.Level, RightBracket)
	result = PaddingRight(result, " ", l.Padding)
	return result
}

// ======================================================================================

// LogNameLayout name layout
type LogNameLayout struct {
}

// NewLogNameLayout new layout
func NewLogNameLayout() LogNameLayout {
	return LogNameLayout{}
}

//
func (l LogNameLayout) String(message LogMessage) string {
	return fmt.Sprintf("%s%s%s", LeftBracket, message.Name, RightBracket)
}

// ======================================================================================

// LogIndentLayout indent layout
type LogIndentLayout struct {
}

// NewLogIndentLayout new layout
func NewLogIndentLayout() LogIndentLayout {
	return LogIndentLayout{}
}
func (l LogIndentLayout) String(message LogMessage) string {
	formatter := Formatter{}
	return fmt.Sprintf("%s", formatter.Compose("\t", message.Indent))
}

// ======================================================================================

// LogMessageLayout message layout
type LogMessageLayout struct {
}

// NewLogMessageLayout new layout
func NewLogMessageLayout() LogMessageLayout {
	return LogMessageLayout{}
}
func (l LogMessageLayout) String(message LogMessage) string {
	formatter := Formatter{}
	content := formatter.ToString(message.Message, formatter.ToString(message.Data...))
	return fmt.Sprintf("%s", content)
}

// ======================================================================================

// LogComposeLayout a special layout
type LogComposeLayout struct {
	Layouts []ILogLayout
}

// NewComposeLayout create compose layout
func NewComposeLayout() *LogComposeLayout {
	return &LogComposeLayout{}
}

// AddLayouts add layouts
func (l *LogComposeLayout) AddLayouts(layouts ...ILogLayout) *LogComposeLayout {
	for _, item := range layouts {
		l.Layouts = append(l.Layouts, item)
	}
	return l
}

func (l *LogComposeLayout) String(message LogMessage) string {
	builder := NewStringBuilder()
	separator := " "

	for _, layout := range l.Layouts {
		builder.Append(layout.String(message)).Append(separator)
	}

	return builder.String()
}

// ======================================================================================

// LayoutCreator creator
type LayoutCreator func() ILogLayout

// LayoutRepository respository for layout creator
type LayoutRepository struct {
	Layouts map[string]LayoutCreator
}

// NewLayoutRepository create new repository
func NewLayoutRepository() LayoutRepository {
	repository := LayoutRepository{make(map[string]LayoutCreator)}

	nameLayoutCreator := func() ILogLayout {
		return NewLogNameLayout()
	}
	timeLayoutCreator := func() ILogLayout {
		return NewLogTimeLayout()
	}
	levelLayoutCreator := func() ILogLayout {
		return NewLogLevelLayout()
	}
	indentLayoutCreator := func() ILogLayout {
		return NewLogIndentLayout()
	}
	messageLayoutCreator := func() ILogLayout {
		return NewLogMessageLayout()
	}
	composeLayoutCreator := func() ILogLayout {
		return NewComposeLayout()
	}

	repository.Layouts[Name] = nameLayoutCreator
	repository.Layouts[NameLayout] = nameLayoutCreator
	repository.Layouts[Time] = timeLayoutCreator
	repository.Layouts[TimeLayout] = timeLayoutCreator
	repository.Layouts[Level] = levelLayoutCreator
	repository.Layouts[LevelLayout] = levelLayoutCreator
	repository.Layouts[Indent] = indentLayoutCreator
	repository.Layouts[IndentLayout] = indentLayoutCreator
	repository.Layouts[Message] = messageLayoutCreator
	repository.Layouts[MessageLayout] = messageLayoutCreator
	repository.Layouts[Compose] = composeLayoutCreator
	repository.Layouts[ComposeLayout] = composeLayoutCreator

	return repository
}

// GetLayout get layout instance by name
func (r LayoutRepository) GetLayout(name string) ILogLayout {
	creator, ok := r.Layouts[name]
	if !ok {
		return nil
	}
	return creator()
}

// BuildLayout build a layout with given layout names
func (r LayoutRepository) BuildLayout(names ...string) ILogLayout {
	layout := NewComposeLayout()

	for _, name := range names {
		item := r.GetLayout(name)
		if item == nil {
			continue
		}

		layout.AddLayouts(item)
	}
	return layout
}
