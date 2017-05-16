package utils

import (
	. "gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestStringBuffer(t *testing.T) {
	Equal(t, NewStringBuffer("test").Join("!").String(), "test!")
	str := NewStringBuffer("").Join("Hello").Join(",").Join([]byte{' ', 'w', 'o', 'r', 'l', 'd'}).Join("!").String()
	Equal(t, str, "Hello, world!")
	Equal(t, NewStringBuffer("").Join("Hello").Join(",").Join([]byte{' ', 'w', 'o', 'r', 'l', 'd'}).Join("!").String(), "Hello, world!")
}
