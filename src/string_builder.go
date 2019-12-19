package logging

import (
	"strings"
)

// StringBuilder string builder wraps the strings.Builder from go library
type StringBuilder struct {
	internalBuilder *strings.Builder
	Lines           int
}

// NewStringBuilder create new string builder
func NewStringBuilder() *StringBuilder {
	var rawBuilder strings.Builder
	builder := new(StringBuilder)
	builder.Lines = 0
	builder.internalBuilder = &rawBuilder
	return builder
}

// Append append string
func (b *StringBuilder) Append(value string) *StringBuilder {
	b.internalBuilder.WriteString(value)
	b.Lines = 1
	return b
}

// Reset reset to empty
func (b *StringBuilder) Reset() *StringBuilder {
	b.internalBuilder.Reset()
	b.Lines = 0
	// var rawBuilder strings.Builder
	// // builder := new(StringBuilder)
	// b.internalBuilder = &rawBuilder
	return b
}

// AppendLine append a line of string
func (b *StringBuilder) AppendLine(value string) *StringBuilder {
	b.internalBuilder.WriteString(value)
	b.internalBuilder.WriteString("\n")
	b.Lines++
	return b
}

// ToString convert string builder to string
func (b *StringBuilder) String() string {
	return b.internalBuilder.String()
}
