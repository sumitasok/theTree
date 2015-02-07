package byteOps

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrepareKey(t *testing.T) {
	assert := assert.New(t)

	assert.Empty(PrepareKey([]byte(``)))

	assert.Equal("abc", PrepareKey([]byte(`  abc} `)))
	assert.Equal("abc", PrepareKey([]byte(`  abc`)))
}

func TestPrepareValueByte(t *testing.T) {
	assert := assert.New(t)

	aVal := []byte(`" \"value"`)
	eVal := []byte(` \"value`)

	assert.Equal(eVal, PrepareValueByte(aVal))

	aVal = []byte(` "\"value\" " `)
	eVal = []byte(`\"value\" `)

	assert.Equal(eVal, PrepareValueByte(aVal))

	aVal = []byte(`value\"" `)
	eVal = []byte(`value\`)

	assert.Equal(eVal, PrepareValueByte(aVal))

	aVal = []byte(`value\" `)
	eVal = []byte(`value\`)

	assert.Equal(eVal, PrepareValueByte(aVal))
}
