package bin

import (
	"testing"

	"github.com/zeebo/assert"
)

func TestTake(t *testing.T) {

	bytes := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
	remaining, value, err := Take(bytes, 3)
	assert.NoError(t, err)
	assert.Equal(t, remaining, []byte{0x04, 0x05})
	assert.Equal(t, value, []byte{0x01, 0x02, 0x03})

}
