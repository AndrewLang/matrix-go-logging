package logging

import (
	"fmt"
	"strings"
	"time"
)

// Formatter used to format text/string
type Formatter struct {
}

// NewFormatter create new formatter
func NewFormatter() *Formatter {
	return &Formatter{}
}

// Format give parameters
func (f Formatter) Format(format string, param ...interface{}) string {
	return fmt.Sprintf(format, param...)
}

// FormatConsoleColor format text with console color format
func (f Formatter) FormatConsoleColor(value string, color string) string {
	return fmt.Sprintf("%s%s%s", color, value, ResetAllStyle)
}

// FormatConsoleBgColor format text with console color format
func (f Formatter) FormatConsoleBgColor(value string, color string) string {
	return fmt.Sprintf("%s%s%s", color, value, ResetAllStyle)
}

// FormatConsoleWith256Color Format text with console 256 color format
func (f Formatter) FormatConsoleWith256Color(value string, num int) string {
	return fmt.Sprintf("%s%d%s%s%s", Color256Start, num, Color256End, value, ResetAllStyle)
}

// FormatConsoleBgWith256Color Format background with console 256 color format
func (f Formatter) FormatConsoleBgWith256Color(value string, num int) string {
	return fmt.Sprintf("%s%d%s%s%s", Color256BgStart, num, Color256End, value, ResetAllStyle)
}

// FormatError format
func (f Formatter) FormatError(err error) string {
	return fmt.Sprintf("%s", err)
}

// FormatDateTime format given date
func (f Formatter) FormatDateTime(date time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d.%02d", date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond()/1000000)
}

// FormatNow format now
func (f Formatter) FormatNow() string {
	return f.FormatDateTime(time.Now())
}

// Compose compose string with give char and count
func (f Formatter) Compose(char string, count int) string {
	builder := NewStringBuilder()
	for i := 0; i < count; i++ {
		builder.Append(char)
	}
	return builder.String()
}

// ToString conver to string
func (f Formatter) ToString(objects ...interface{}) string {
	var builder strings.Builder
	spearator := " "
	length := len(objects)

	for index, value := range objects {

		switch valueType := value.(type) {
		case fmt.Stringer:
			builder.WriteString(valueType.String())
		case string:
			builder.WriteString(valueType)
		default:
			builder.WriteString(fmt.Sprintf("%v", valueType))
		}

		if index < length-1 {
			builder.WriteString(spearator)
		}
	}
	return builder.String()
}
