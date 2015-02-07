package byteOps

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrepareKey(t *testing.T) {
	assert := assert.New(t)

	assert.Empty(prepareKey([]byte(``)))

	assert.Equal("abc", prepareKey([]byte(`  abc} `)))
	assert.Equal("abc", prepareKey([]byte(`  abc`)))
}
