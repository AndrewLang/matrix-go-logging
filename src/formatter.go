package logging

import (
	"fmt"
	"strings"
	"time"
)

// Formatter used to
type Formatter struct {
}

// Format give parameters
func (f Formatter) Format(format string, param ...interface{}) string {
	return fmt.Sprintf(format, param...)
}

// FormatError format
func (f Formatter) FormatError(err error) string {
	return fmt.Sprintf("%s", err)
}

// FormatDateTime format
func (f Formatter) FormatDateTime(date time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d.%02d", date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond()/1000000)
}

// FormatNow format
func (f Formatter) FormatNow() string {
	return f.FormatDateTime(time.Now())
}

// Compose compose string with give char and count
func (f Formatter) Compose(char string, count int) string {
	// fmt.Println("Compose string with: ", char, count)
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
