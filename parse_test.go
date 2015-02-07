package theTree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringParse(t *testing.T) {
	assert := assert.New(t)

	byteArr := []byte(`{key: "value \"", "key2" : "value_2"}`)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)

	err := SetNodeValue(node, byteArr)
	assert.NoError(err)

	fmt.Println("-----------------------------------")

	for i, nodeC := range node.Children {
		str, ok := nodeC.Value.(string)
		fmt.Println("nodeC", i, nodeC.Key, "Value", ok, str)
	}
}
