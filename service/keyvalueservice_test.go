package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {

	store := Store()
	store.Put("1", "1")

	assert := assert.New(t)
	assert.Equal("1", store.data["1"])
	assert.Equal("", store.data[""])
}

func TestGet(t *testing.T) {

	store := Store()
	store.Put("1", "1")

	assert := assert.New(t)
	assert.Equal("1", store.Get("1"))
	assert.Equal("", store.Get("2"))
}
