package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

/* Remember about naming
Go naming convention
The convention in Go is to use MixedCaps or mixedCaps (simply camelCase) rather than underscores to write multi-word names.
If an identifier needs to be visible outside the package, its first character should be uppercase.
If you don't have the intention to use it in another package, you can safely stick to mixedCaps
*/

// Block - the block for blockChain
type Block struct {
	Timestamp int64
	Data      []byte
	Hash      []byte
	PrevBlock []byte
}

func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlock, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// Creating new block of chain
func newBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.setHash()
	return block
}
