package theTree

import (
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
	assert.Error(err)
	assert.Equal(err_key_already_exists, err.Error())

}

func TestUpdateChild(t *testing.T) {
	assert := assert.New(t)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)
	node.Set("value")

	newNode, err := node.Append("key")
	newNode.Set("value")

	newNode, err = node.UpdateChild("key", 123)
	assert.Equal("key", node.Children[0].Key)
	assert.Equal(123, node.Children[0].Value)
	assert.NoError(err)

	newNode, err = node.UpdateChild("nonexistantkey", 123)
	assert.Error(err)
	assert.Equal(err_key_doesnt_exist, err.Error())
	assert.Empty(newNode.Key)
	assert.Empty(newNode.Value)

}

func TestChild(t *testing.T) {
	assert := assert.New(t)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)
	node.Set("value")

	newNode, _ := node.Append("number")
	newNode.Set(123)

	newNode, _ = node.Append("word")
	newNode.Set("love")

	assert.Equal(2, len(node.Children))

	found, errFound := node.Child("word")
	assert.NoError(errFound)
	assert.Equal("love", found.Value)

	notFound, errNotFound := node.Child("mock")
	assert.Error(errNotFound)
	assert.NotEqual("love", notFound.Value)
}

func TestFind(t *testing.T) {
	assert := assert.New(t)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)
	node.Set("value")

	nodeA, _ := node.Append("number")
	nodeA.Set(123)

	nodeB, _ := nodeA.Append("key")
	nodeB.Set("rocket")

	_, errNonExistant := node.Find("love:number:key")
	assert.Error(errNonExistant)

	actNode, actErr := node.Find("root")
	assert.Equal(node.Value, actNode.Value)
	assert.NoError(actErr)

	actNode, actErr = node.Find("root:number")
	assert.Equal(nodeA.Value, actNode.Value)
	assert.NoError(actErr)

	actNode, actErr = node.Find("root:number:key")
	assert.Equal(nodeB.Value, actNode.Value)
	assert.NoError(actErr)
}

func TestAncestry(t *testing.T) {
	assert := assert.New(t)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)
	node.Set("value")

	nodeA, _ := node.Append("number")
	nodeA.Set(123)

	nodeB, _ := nodeA.Append("key")
	nodeB.Set("rocket")

	assert.Equal("root:number:key", nodeB.Ancestry())
}

func TestRoot(t *testing.T) {
	assert := assert.New(t)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)
	node.Set("value")

	nodeA, _ := node.Append("number")
	nodeA.Set(123)

	nodeB, _ := nodeA.Append("key")
	nodeB.Set("rocket")

	assert.Equal(node, nodeB.Root())
}

func TestCount(t *testing.T) {
	assert := assert.New(t)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)
	node.Set("value")

	nodeA, _ := node.Append("number")
	nodeA.Set(123)

	nodeB, _ := nodeA.Append("key")
	nodeB.Set("rocket")

	assert.Equal(1, node.Count())
}

func TestCountPre(t *testing.T) {
	assert := assert.New(t)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)
	node.Set("value")

	nodeA, _ := node.Append("number")
	nodeA.Set(123)

	nodeB, _ := nodeA.Append("key")
	nodeB.Set("rocket")

	assert.Equal(0, node.CountPre())
	assert.Equal(1, nodeA.CountPre())
	assert.Equal(2, nodeB.CountPre())
}

func TestCountDeep(t *testing.T) {
	assert := assert.New(t)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)
	node.Set("value")

	nodeA, _ := node.Append("number")
	nodeA.Set(123)

	nodeB, _ := nodeA.Append("key")
	nodeB.Set("rocket")

	nodeC, _ := nodeA.Append("key 2")
	nodeC.Set("rocket 2")

	nodeD, _ := nodeB.Append("D")
	nodeD.Set("rocket")

	assert.Equal(4, node.CountDeep())
	assert.Equal(3, nodeA.CountDeep())
	assert.Equal(1, nodeB.CountDeep())
}

func TestJson(t *testing.T) {
	assert := assert.New(t)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)
	node.Set("value")

	nodeA, _ := node.Append("number")
	nodeA.Set(123)

	nodeB, _ := nodeA.Append("key")
	nodeB.Set("rocket")

	_, err := node.Json()
	assert.NoError(err)
}

/*
func TestCount(t *testing.T) {
	assert := assert.New(t)
}
*/
