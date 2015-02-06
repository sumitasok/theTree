package theTree

// A double quote ends with another double quotes
//  - escape \"
// open { ends in } where any of } occurring in " or '
// is escaped

// Node{}
// set value as the string
// then parse and create CHild Nodes and assign their values

import (
	// "bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringParse(t *testing.T) {
	// assert := assert.New(t)

	// byteArray := []byte(`{key: "value \"", "key2" : []}`)

	// engine := Normal{}
	// key := "root"

	// node := Init(engine, key)

	// SetNodeValue(node, byteArray)

	fmt.Println("-----------------------------------")
}

func TestBytePluckByte(t *testing.T) {
	assert := assert.New(t)

	byteArr := bytePluckByte([]byte(`  {  `), R_SPACE)
	assert.Equal(string([]byte(`{`)), string(byteArr))
}

func TestByteRemoveByte(t *testing.T) {
	assert := assert.New(t)

	byteArr := []byte(``)
	assert.Empty(byteArr)

	byteArr = byteRemoveByte([]byte(` `), R_SPACE)
	assert.Empty(byteArr)

	byteArr = byteRemoveByte([]byte(`  `), R_SPACE)
	assert.Empty(byteArr)

	byteArr = byteRemoveByte([]byte(`  {`), R_SPACE)
	assert.Equal([]byte(`{`), byteArr)

}

func TestByteRemoveByteFromBack(t *testing.T) {
	assert := assert.New(t)

	byteArr := []byte(``)
	assert.Empty(byteArr)

	byteArr = byteRemoveByteFromBack([]byte(` `), R_SPACE)
	assert.Empty(byteArr)

	byteArr = byteRemoveByteFromBack([]byte(`  `), R_SPACE)
	assert.Empty(byteArr)

	byteArr = byteRemoveByteFromBack([]byte(`}  `), R_SPACE)
	assert.Equal(string([]byte(`}`)), string(byteArr))

}

func TestByteIs(t *testing.T) {
	assert := assert.New(t)

	result := byteIs([]byte(` `), R_OPEN_CURL)
	assert.False(result)

	result = byteIs([]byte(` {`), R_OPEN_CURL)
	assert.True(result)

	result = byteIs([]byte(` "`), R_OPEN_CURL)
	assert.False(result)
}

// ----------------------------------------------------

func TestReverse(t *testing.T) {
	assert := assert.New(t)

	str := []byte(`  {]  `)
	strRev := []byte(`  ]{  `)

	assert.Equal(string(strRev), string(Reverse(str)))

	str = []byte(`  {][  `)
	strRev = []byte(`  []{  `)

	assert.Equal(string(strRev), string(Reverse(str)))

}

func TestByteIsHash(t *testing.T) {
	assert := assert.New(t)

	result := byteIsHash([]byte(` `))
	assert.False(result)

	result = byteIsHash([]byte(` {`))
	assert.True(result)

	result = byteIsHash([]byte(` "`))
	assert.False(result)
}

func TestByteRemoveSpace(t *testing.T) {
	assert := assert.New(t)

	byteArr := []byte(``)
	assert.Empty(byteArr)

	byteArr = byteRemoveSpace([]byte(` `))
	assert.Empty(byteArr)

	byteArr = byteRemoveSpace([]byte(`  `))
	assert.Empty(byteArr)

	byteArr = byteRemoveSpace([]byte(`  {`))
	assert.Equal([]byte(`{`), byteArr)

}

// for i, r := range bytes.Runes(byteArray) {
// 	fmt.Println(i, " - ", r, " - ", string(r))
// }