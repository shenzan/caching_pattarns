package better

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"hash"
)

var (
	hsha1 = sha1.New()
	hsha2 = sha256.New()
	hmd5  = md5.New()
)

type filter struct {
	arr [1000]bool
}

func (f *filter) Add(s string) {
	space := len(f.arr)
	f.arr[hashPosition(hsha1, s, space)] = true
	f.arr[hashPosition(hsha2, s, space)] = true
	f.arr[hashPosition(hmd5, s, space)] = true
}

func (f *filter) Exist(s string) bool {
	space := len(f.arr)
	return f.arr[hashPosition(hsha1, s, space)] &&
		f.arr[hashPosition(hsha2, s, space)] &&
		f.arr[hashPosition(hmd5, s, space)]
}

func hashPosition(h hash.Hash, s string, space int) int {
	hs := createHash(h, s)
	if hs < 0 {
		hs = -hs
	}
	return hs % space
}

func createHash(h hash.Hash, input string) int {
	bits := h.Sum([]byte(input))
	buffer := bytes.NewBuffer(bits)
	result, _ := binary.ReadVarint(buffer)
	return int(result)
}
