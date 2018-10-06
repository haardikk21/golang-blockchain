package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Timestamp     int64          // Block creation timestamp
	Transactions  []*Transaction // All transactions in the block
	PrevBlockHash []byte         // Hash of the previous block
	Hash          []byte         // Hash of this block
	Nonce         int            // Nonce of this block
}

// NewBlock creates a new Block type with provided data and prevBlockHash,
// It calls SetHash and returns the newly created Block
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// HashTransactions hashes all the transactions in a block
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}
