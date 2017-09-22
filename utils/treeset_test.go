package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestTreeSet(t *testing.T) {
	treeSet := NewTreeSet()
	err := treeSet.Set("/", 0)
	assert.Equal(t, err, nil)
	err = treeSet.Set("/0", 1)
	assert.Equal(t, err, nil)
	err = treeSet.Set("/a", 2)
	assert.Equal(t, err, nil)
	err = treeSet.Set("/a/b", 3)
	assert.Equal(t, err, nil)
	err = treeSet.Set("/a/b", 6)
	assert.Equal(t, err, nil)
	err = treeSet.Set("/a/b/c", 9)
	assert.Equal(t, err, nil)
	err = treeSet.Set("/o/b", 4)
	values, err := treeSet.Get("/")
	assert.Equal(t, err, nil)
	fmt.Println(values)
}