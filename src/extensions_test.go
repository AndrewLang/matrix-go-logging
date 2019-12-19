package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	actual := Join("hello ", "world")

	assert.Equal(t, actual, "hello world", "Join result should be 'hello world'")
}

func TestJoinWith(t *testing.T) {
	actual := JoinWith(" ", "hello", "world")

	assert.Equal(t, actual, "hello world", "Join result should be 'hello world'")
}

func TestJoinStrings(t *testing.T) {
	actual := JoinStrings([]string{"hello", "world"}, " ")

	assert.Equal(t, actual, "hello world", "Join result should be 'hello world'")
}
