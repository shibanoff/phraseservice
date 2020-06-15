// Package hashcode provides simple way to get hashcode of string.
package hashcode

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"strconv"
)

// HashCode is the interface implemented by an object that can load string and return hash of string as integer.
// On same input data it must return the same integer value.
type HashCode interface {
	// LoadString writes string to the hasher
	LoadString(str string)
	// GetCode returns lower 8 bytes from hash of input string as int64
	GetCode() int64
}

type hashCode struct {
	hasher hash.Hash
}

// New returns a new HashCode with given hash.
func New(hasher hash.Hash) HashCode {
	return hashCode{hasher: hasher}
}

// NewSHA265 returns a new HashCode with SHA256 hash.
func NewSHA265() HashCode {
	return New(sha256.New())
}

func (h hashCode) LoadString(str string) {
	h.hasher.Reset()
	h.hasher.Write([]byte(str))
}

func (h hashCode) GetCode() int64 {
	hashAsBytes := h.hasher.Sum(nil)
	return fitHashIntoInt(hashAsBytes)
}

func fitHashIntoInt(hash []byte) int64 {
	if len(hash) == 0 {
		return 0
	}

	if len(hash) >= 8 {
		// MaxInt64 = 9223372036854775807 = 7FFFFFFFFFFFFFFF
		// Get 8 lower bytes
		hash = hash[len(hash)-8:]
		// and cut off higher bit.
		hash[0] &= 0x7F
	}

	i, err := strconv.ParseInt(hex.EncodeToString(hash), 16, 64)
	if err != nil {
		panic(err)
	}

	return i
}
