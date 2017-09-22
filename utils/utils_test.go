package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringBuffer(t *testing.T) {
	assert.Equal(t, NewStringBuffer("test").Join("!").String(), "test!")
	str := NewStringBuffer("").Join("Hello").Join(",").Join([]byte{' ', 'w', 'o', 'r', 'l', 'd'}).Join("!").String()
	assert.Equal(t, str, "Hello, world!")
	assert.Equal(t, NewStringBuffer("").Join("Hello").Join(",").Join([]byte{' ', 'w', 'o', 'r', 'l', 'd'}).Join("!").String(), "Hello, world!")
}
