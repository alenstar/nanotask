package config

import (
	. "gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestConfig(t *testing.T) {
	NotEqual(t, String("key", ""), "hello")

	Equal(t, String("key1", ""), "hello")
	Equal(t, String("key2", ""), "world")

	Equal(t, Int("keyInt1"), 3)
	Equal(t, Int("keyInt2"), 4)

	Equal(t, Bool("keyBool1"), false)
	Equal(t, Bool("keyBool2"), true)

	Equal(t, Float("keyFloat1"), 2.71828)
	Equal(t, Float("keyFloat2"), 3.1415926)
}
