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

func TestFetchInnerCurls(t *testing.T) {
	assert := assert.New(t)

	str := string([]byte(`{"holder": "holds"}`))
	result := fetchInnerCurls(str)
	assert.Equal(string([]byte(`"holder": "holds"`)), result.Content)

	str = string([]byte(`{"holder": "holds","hash": {"inner": "peace"}}`))
	result = fetchInnerCurls(str)
	assert.Equal(string([]byte(`"holder": "holds"`)), result.Next[0].Content)
	// assert.Equal(string([]byte(`"inner": "peace"`)), result.Next[1].Next[0].Content)

}

func TestOnlyColon(t *testing.T) {
	assert := assert.New(t)

	str := string([]byte(`"holder": "holds"`))
	e := Element{Content: str}
	e.OnlyColon()
	assert.Equal("holder", e.Key)
	assert.Equal("holds", e.Value)
}
