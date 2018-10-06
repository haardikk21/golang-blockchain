package main

import (
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Timestamp     int64  // Block creation timestamp
	Data          []byte // Actual valuable information being contained in the Block
	PrevBlockHash []byte // Hash of the previous block
	Hash          []byte // Hash of this block
	Nonce         int    // Nonce of this block
}

// NewBlock creates a new Block type with provided data and prevBlockHash,
// It calls SetHash and returns the newly created Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}
