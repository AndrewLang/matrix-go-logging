package logging

import (
	"fmt"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	formatter := Formatter{}
	err := NewException("this is an exception")

	actual := formatter.ToString("this is test", 1, 2, true, err)
	fmt.Println(actual)

	// assert.Equal(t, actual, "hello world", "Join result should be 'hello world'")
}
