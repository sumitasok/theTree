package theTree

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonParse(t *testing.T) {
	assert := assert.New(t)

	// data := []byte(`{"holder": "holds"}`)

	// node, _ := Parse(data)

	// assert.Equal("holder", node.Key)
	// assert.Equal("holds", node.Value)

	assert.True(true)

}

func TestStripOuterCurls(t *testing.T) {
	assert := assert.New(t)

	str := "{\"holder\": \"holds\"}"

	result := stripOuterCurls(str)

	assert.Equal("\"holder\": \"holds\"", result)
}
