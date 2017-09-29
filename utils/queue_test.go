package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Push(32)
	// assert.Equal(t, 32, int(q.root.value.(int)))
	q.Push(33)
	q.Push(34)
	q.Push(35)

	assert.Equal(t, 4, int(q.Size()))

	assert.Equal(t, 32, q.Pop().(int))
	assert.Equal(t, 33, q.Pop().(int))
	assert.Equal(t, 34, q.Pop().(int))
	assert.Equal(t, 35, q.Pop().(int))
}
