package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

func New() string{
	var u [16]byte
	_, _ = io.ReadFull(rand.Reader, u[:])

	uHex := make([]byte, 36)
	hex.Encode(uHex[0:8], u[0:4])
	uHex[8] = '-'
	hex.Encode(uHex[9:13], u[4:6])
	uHex[13] = '-'
	hex.Encode(uHex[14:18], u[6:8])
	uHex[18] = '-'
	hex.Encode(uHex[19:23], u[8:10])
	uHex[23] = '-'
	hex.Encode(uHex[24:], u[10:])

	return string(uHex)
}