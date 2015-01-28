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

	assert.NotNil(node)
}

func TestSet(t *testing.T) {
	assert := assert.New(t)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)
	node.Set("value")

	assert.Equal("value", node.Value)
	assert.Equal("string", node.DataType)
}

func TestAppend(t *testing.T) {
	assert := assert.New(t)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)
	node.Set("value")

	newNode, err := node.Append("key")
	assert.Equal("key", node.Children[0].Key)
	assert.Equal("key", newNode.Key)
	assert.NoError(err)

	newNode, err = node.Append("key")
	assert.Equal("key", node.Children[0].Key)
	assert.NotEqual("key", newNode.Key)
	assert.Error(err)
	assert.Equal(err_key_already_exists, err.Error())

}
