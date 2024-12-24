package mymap

import "hash/crc32"

const arraySize = 3

type hashData[T any] struct {
	key string
	val T
}

type HashMap[T any] struct {
	arr [arraySize][]hashData[T]
}

func hashfn(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (h *HashMap[T]) Add(key string, val T) {
	hash := hashfn(key)
	data := hashData[T]{
		key: key,
		val: val,
	}
	h.arr[hash%arraySize] = append(h.arr[hash%arraySize], data)
}

func (h *HashMap[T]) Get(key string) (T, bool) {
	hash := hashfn(key)
	for _, data := range h.arr[hash%arraySize] {
		if data.key == key {
			return data.val, true
		}
	}
	var temp T
	return temp, false
}
