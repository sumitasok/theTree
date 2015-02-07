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

	assert.Equal(printStrOfBytes([]byte(`key`)), printStrOfBytes(kvList[0].Key))
	assert.Equal(printStrOfBytes([]byte(`"value \"`)), printStrOfBytes(kvList[0].Value))

	assert.Equal(printStrOfBytes([]byte(`key2`)), printStrOfBytes(kvList[1].Key))
	assert.Equal(printStrOfBytes([]byte(`"value_2"`)), printStrOfBytes(kvList[1].Value))

	assert.True(true)
}

func TestPrepareValue(t *testing.T) {
	assert := assert.New(t)

	aVal := []byte(`" \"value"`)
	eVal := []byte(` \"value`)

	assert.Equal(eVal, PrepareValue(aVal))

	aVal = []byte(` "\"value\" " `)
	eVal = []byte(`\"value\" `)

	assert.Equal(eVal, PrepareValue(aVal))
}

func printStrOfBytes(byteArr interface{}) string {
	if b, ok := byteArr.([]uint8); ok {
		return string(b)
	} else {
		return "error in printStrOfBytes"
	}
}
