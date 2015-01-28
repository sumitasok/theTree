package theTree

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeInit(t *testing.T) {
	assert := assert.New(t)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)

	assert.True(true)
	assert.Nil(node)
}
