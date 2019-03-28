package cryptonight

// #cgo CFLAGS: -march=ivybridge
// #include "cn.h"
import "C"
import "unsafe"

// SlowHash do cryptonight slow hash function on a byte slice.
func SlowHash(d []byte) []byte {
	l := len(d)
	b := make([]C.char, l)
	for i, c := range d {
		b[i] = C.char(c)
	}
	var cr [32]C.char
	bptr := unsafe.Pointer(&b[0])
	C.cn_slow_hash(bptr, C.size_t(l), &cr[0], 0, 0)
	r := make([]byte, 32)
	for i, c := range cr {
		r[i] = byte(c)
	}
	return r
}
