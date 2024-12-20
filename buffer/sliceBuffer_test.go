package buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	buf := NewSliceBuffer[byte]()

	buf.Write([]byte{1, 2, 3, 4, 5})

	assert.Equal(t, 5, buf.Readable())
}
func TestRead(t *testing.T) {
	buf := NewSliceBuffer[byte]()
	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, buf.Readable())

	readedData := buf.Read(4)
	for idx, val := range readedData {
		assert.Equal(t, byte(idx+1), val)
	}

	assert.Equal(t, 0, buf.Readable())
}

func TestWriteAndRead(t *testing.T) {
	buf := NewSliceBuffer[byte]()
	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, buf.Readable())

	buf.Write([]byte{5, 6})
	assert.Equal(t, 6, buf.Readable())

	readedData := buf.Read(4)
	for idx, val := range readedData {
		assert.Equal(t, byte(idx+1), val)
	}

	assert.Equal(t, 2, buf.Readable())

	buf.Write([]byte{7, 8, 9})
	assert.Equal(t, 5, buf.Readable())

	readedData = buf.Read(4)
	for idx, val := range readedData {
		assert.Equal(t, byte(idx+5), val)
	}
	assert.Equal(t, 1, buf.Readable())
}
