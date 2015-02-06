package theTree

// A double quote ends with another double quotes
//  - escape \"
// open { ends in } where any of } occurring in " or '
// is escaped

// Node{}
// set value as the string
// then parse and create CHild Nodes and assign their values

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// func TestStringParse(t *testing.T) {
// 	assert := assert.New(t)

// 	byteArr := []byte(`{key: "value \"", "key2" : []}`)

// 	engine := Normal{}
// 	key := "root"

// 	node := Init(engine, key)

// 	err := SetNodeValue(node, byteArr)
// 	expectedByteArr := []byte(` "key2" : []`)
// 	assert.NoError(err)

// 	actualByteArr, _ := node.Value.([]uint8)
// 	assert.Equal(string(expectedByteArr), string(actualByteArr))

// 	fmt.Println("-----------------------------------")
// }

func TestParseNode(t *testing.T) {
	assert := assert.New(t)

	byteArr := []byte(`key: "value \"", "key2" : value_2`)

	for i, r := range bytes.Runes(byteArr) {
		fmt.Println(i, " - ", r, " - ", string(r))
	}

	engine := Normal{}
	key := "root"

	node := Init(engine, key)

	node.Set(byteArr)

	parseNode(node, byteArr)

	newNode, err := node.Child("key")
	assert.Equal("key", newNode.Key)
	assert.NoError(err)

	newNodeValue, err1 := newNode.Value.([]uint8)
	assert.Equal(string([]byte(`"value \""`)), string(newNodeValue))
	assert.True(err1)

	fmt.Println("----------------------------")

	node = Init(engine, key)
	node.Set(byteArr)

	parseNode(node, byteArr)

	for i, n := range node.Children {
		fmt.Println(i, n.Key)
	}

	newNode, err = node.Child(string([]byte(`"key2"`)))
	assert.Equal("key2", newNode.Key)
	assert.NoError(err)

	newNodeValue, err1 = newNode.Value.([]uint8)
	assert.Equal(string([]byte(`"value \""`)), string(newNodeValue))
	assert.True(err1)

}
