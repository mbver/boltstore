package boltstore

import (
	"bytes"
	"encoding/binary"

	"github.com/hashicorp/go-msgpack/codec"
)

// Converts bytes to an integer
func toUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// Converts a uint to a byte slice
func toBytes(u uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u)
	return buf
}
func decode(buf []byte, out interface{}) error {
	r := bytes.NewBuffer(buf)
	hd := codec.MsgpackHandle{}
	dec := codec.NewDecoder(r, &hd)
	return dec.Decode(out)
}

// Encode writes an encoded object to a new bytes buffer
func encode(in interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	hd := codec.MsgpackHandle{}
	enc := codec.NewEncoder(buf, &hd)
	err := enc.Encode(in)
	return buf.Bytes(), err
}
