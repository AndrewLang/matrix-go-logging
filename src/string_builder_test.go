package logging

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringBuilderAppend(t *testing.T) {
	builder := NewStringBuilder()
	builder.Append("hello ").Append("world")
	actual := builder.String()

	assert.Equal(t, actual, "hello world", "StringBuilder result should be 'hello world'")
}

func TestStringBuilderAppendLine(t *testing.T) {
	builder := NewStringBuilder()
	builder.AppendLine("hello ").AppendLine("world")
	actual := builder.String()

	fmt.Println("AppendLine resutl: ", actual)
	assert.Equal(t, actual, "hello \nworld\n", "StringBuilder result should be 'hello world'")
}

func TestStringBuilderReset(t *testing.T) {
	builder := NewStringBuilder()
	builder.Append("hello ").Append("world")
	actual := builder.String()

	assert.Equal(t, actual, "hello world", "StringBuilder result should be 'hello world'")

	builder.Reset()
	actual = builder.String()
	assert.Equal(t, "", actual, "After reset value is not empty")
}
