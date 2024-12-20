package buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRingWrite(t *testing.T) {
	buf := NewRingBuffer[byte](10)

	buf.Write([]byte{1, 2, 3, 4})

	assert.Equal(t, 4, buf.Readable())
}
func TestRingRead(t *testing.T) {
	buf := NewSliceBuffer[byte]()
	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, buf.Readable())

	readedData := buf.Read(4)
	for idx, val := range readedData {
		assert.Equal(t, byte(idx+1), val)
	}

	assert.Equal(t, 0, buf.Readable())
}

func TestRingWriteAndRead(t *testing.T) {
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

func TestRingOverWrite(t *testing.T) {
	buf := NewRingBuffer[byte](5)
	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, buf.Readable())
	assert.Equal(t, 1, buf.Writeable())

	buf.Write([]byte{5})
	assert.Equal(t, 5, buf.Readable())
	assert.Equal(t, 0, buf.Writeable())

	writed := buf.Write([]byte{6})
	assert.Equal(t, 0, writed)
	assert.Equal(t, 5, buf.Readable())
	assert.Equal(t, 0, buf.Writeable())

	readedData := buf.Read(4)
	for idx, val := range readedData {
		assert.Equal(t, byte(idx+1), val)
	}

	assert.Equal(t, 1, buf.Readable())
	assert.Equal(t, 4, buf.Writeable())

	writed = buf.Write([]byte{6, 7, 8})
	assert.Equal(t, 3, writed)
	assert.Equal(t, 3, buf.writePoint)
	assert.Equal(t, 4, buf.Readable())
	assert.Equal(t, 1, buf.Writeable())

	writed = buf.Write([]byte{6, 7, 8})
	assert.Equal(t, 1, writed)
	assert.Equal(t, 5, buf.Readable())
	assert.Equal(t, 0, buf.Writeable())

	readedData = buf.Read(4)
	t.Log(readedData)
	assert.Equal(t, 4, len(readedData))
	assert.Equal(t, 1, buf.Readable())
	assert.Equal(t, 4, buf.Writeable())
	t.Log(buf)
}
