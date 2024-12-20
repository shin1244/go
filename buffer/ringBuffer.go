package buffer

type RingBuffer[T any] struct {
	buf        []T
	readPoint  int
	writePoint int
	isFull     bool
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	buf := &RingBuffer[T]{buf: make([]T, size)}
	return buf
}

func (r *RingBuffer[T]) Write(data []T) int {
	if len(data) == 0 || r.Writeable() == 0 {
		return 0
	}
	var writed int
	if r.writePoint >= r.readPoint {
		writeableToEnd := len(r.buf) - r.writePoint
		writed = min(len(data), writeableToEnd)
	} else {
		writed = min(len(data), r.Writeable())
	}
	copy(r.buf[r.writePoint:], data[:writed])
	r.writePoint += writed
	r.writePoint %= len(r.buf)

	if writed > 0 && r.readPoint == r.writePoint {
		r.isFull = true
	}

	remain := len(data) - writed
	if remain > 0 && r.Writeable() > 0 {
		writed += r.Write(data[writed:])
	}

	return writed
}

func (r *RingBuffer[T]) Writeable() int {
	if r.isFull {
		return 0
	}
	if r.writePoint >= r.readPoint {
		return len(r.buf) - r.writePoint + r.readPoint
	}
	return r.readPoint - r.writePoint
}
func (r *RingBuffer[T]) Readable() int {
	return len(r.buf) - r.Writeable()
}

func (r *RingBuffer[T]) Read(count int) []T {
	if r.Readable() == 0 && count <= 0 {
		return []T{}
	}
	readCnt := min(count, r.Readable())
	rst := make([]T, readCnt)
	if r.readPoint+readCnt >= len(r.buf) {
		copy(rst, r.buf[r.readPoint:])
		copy(rst[len(r.buf)-r.readPoint:], r.buf)
		r.readPoint = (r.readPoint + readCnt) % len(r.buf)
	} else {
		copy(rst, r.buf[r.readPoint:r.readPoint+readCnt])
		r.readPoint += readCnt
	}
	r.isFull = false

	return rst
}
