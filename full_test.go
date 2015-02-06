package theTree

// A double quote ends with another double quotes
//  - escape \"
// open { ends in } where any of } occurring in " or '
// is escaped

// Node{}
// set value as the string
// then parse and create CHild Nodes and assign their values

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringParse(t *testing.T) {
	assert := assert.New(t)

	byteArr := []byte(`{key: "value \"", "key2" : []}`)

	engine := Normal{}
	key := "root"

	node := Init(engine, key)

	err := SetNodeValue(node, byteArr)
	expectedByteArr := []byte(`key: "value \"", "key2" : []`)
	assert.NoError(err)

	actualByteArr, _ := node.Value.([]uint8)
	assert.Equal(string(expectedByteArr), string(actualByteArr))

	fmt.Println("-----------------------------------")
}
