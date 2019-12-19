package logging

import (
	"strings"
)

// Join concate given strings
func Join(values ...string) string {
	var builder strings.Builder

	for _, value := range values {
		builder.WriteString(value)
	}

	return builder.String()
}

// JoinStrings join
func JoinStrings(values []string, spearator string) string {
	var builder strings.Builder

	for index, value := range values {
		builder.WriteString(value)
		if index < len(values)-1 {
			builder.WriteString(spearator)
		}
	}

	return builder.String()
}

// JoinWith join
func JoinWith(spearator string, values ...string) string {
	var builder strings.Builder

	for index, value := range values {
		builder.WriteString(value)
		if index < len(values)-1 {
			builder.WriteString(spearator)
		}
	}

	return builder.String()
}

// PaddingRight padding right
func PaddingRight(original string, place string, maxLen int) string {
	builder := NewStringBuilder()
	formatter := Formatter{}
	length := len(original)

	builder.Append(original)
	if length < maxLen {
		builder.Append(formatter.Compose(place, maxLen-length))
	}

	return builder.String()
}
