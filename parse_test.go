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
		str, ok := nodeC.Value.([]uint8)
		fmt.Println("nodeC", i, nodeC.Key, "Value", ok, string(str))
	}
}

func TestParseKeyValue(t *testing.T) {
	assert := assert.New(t)

	byteArr := []byte(`key: "value \"", "key2" : "value_2"`)

	kvList := parseKeyValue(byteArr)

	assert.Equal(2, len(kvList))

	assert.Equal([]byte(`key`), kvList[0].Key)
	assert.Equal([]byte(`value \`), kvList[0].Value)

	assert.Equal([]byte(`key2`), kvList[1].Key)
	assert.Equal([]byte(`value_2`), kvList[1].Value)

	assert.True(true)
}

func printStrOfBytes(byteArr interface{}) string {
	if b, ok := byteArr.([]uint8); ok {
		return string(b)
	} else {
		return "error in printStrOfBytes"
	}
}
