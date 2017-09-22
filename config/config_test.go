package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert.NotEqual(t, String("key", ""), "hello")

	assert.Equal(t, String("key1", ""), "hello")
	assert.Equal(t, String("key2", ""), "world")

	assert.Equal(t, Int("keyInt1"), 3)
	assert.Equal(t, Int("keyInt2"), 4)

	assert.Equal(t, Bool("keyBool1"), false)
	assert.Equal(t, Bool("keyBool2"), true)

	assert.Equal(t, Float("keyFloat1"), 2.71828)
	assert.Equal(t, Float("keyFloat2"), 3.1415926)
}
