package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// IdentKeyPrefix is the prefix to retrieve all Ident
	IdentKeyPrefix = "Ident/value/"
)

// IdentKey returns the store key to retrieve a Ident from the index fields
func IdentKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
