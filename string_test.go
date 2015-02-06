package theTree

import (
	// "bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByteStripOuterCurls(t *testing.T) {
	assert := assert.New(t)
	byteArr := []byte(`{key: "value \"", "key2" : []}`)

	expectedByteArr := []byte(`key: "value \"", "key2" : []`)

	actualByteArr, err := byteStripOuterCurls(byteArr)
	assert.Equal(string(expectedByteArr), string(actualByteArr))
	assert.NoError(err)

	byteArr = []byte(`key: "value \"", "key2" : []}`)
	expectedByteArr = []byte(`key: "value \"", "key2" : []}`)
	actualByteArr, err = byteStripOuterCurls(byteArr)
	assert.Equal(string(expectedByteArr), string(actualByteArr))
	assert.Error(err)

	byteArr = []byte(`{key: "value \"", "key2" : []`)
	expectedByteArr = []byte(`key: "value \"", "key2" : []`)
	actualByteArr, err = byteStripOuterCurls(byteArr)
	assert.Equal(string(expectedByteArr), string(actualByteArr))
	assert.Error(err)

}

func TestByteRemove(t *testing.T) {
	assert := assert.New(t)

	byteArr, _ := byteRemove([]byte(`{  `), R_OPEN_CURL)
	assert.Equal(string([]byte(`  `)), string(byteArr))

	byteArr, err := byteRemove([]byte(`[  `), R_OPEN_CURL)
	assert.Equal(string([]byte(`[  `)), string(byteArr))
	assert.Error(err)

}

func TestBytePluckByteRecursively(t *testing.T) {
	assert := assert.New(t)

	byteArr := bytePluckByteRecursively([]byte(`  {  `), R_SPACE, R_SPACE)
	assert.Equal(string([]byte(`{`)), string(byteArr))
}

func TestByteRemoveByteRecursively(t *testing.T) {
	assert := assert.New(t)

	byteArr := []byte(``)
	assert.Empty(byteArr)

	byteArr = byteRemoveByteRecursively([]byte(` `), R_SPACE)
	assert.Empty(byteArr)

	byteArr = byteRemoveByteRecursively([]byte(`  `), R_SPACE)
	assert.Empty(byteArr)

	byteArr = byteRemoveByteRecursively([]byte(`  {`), R_SPACE)
	assert.Equal([]byte(`{`), byteArr)

}

func TestByteRemoveByteRecursivelyFromBack(t *testing.T) {
	assert := assert.New(t)

	byteArr := []byte(``)
	assert.Empty(byteArr)

	byteArr = byteRemoveByteRecursivelyFromBack([]byte(` `), R_SPACE)
	assert.Empty(byteArr)

	byteArr = byteRemoveByteRecursivelyFromBack([]byte(`  `), R_SPACE)
	assert.Empty(byteArr)

	byteArr = byteRemoveByteRecursivelyFromBack([]byte(`}  `), R_SPACE)
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
