package baseutils

import "bytes"

func JoinBytes(sep []byte, pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, sep)
}
