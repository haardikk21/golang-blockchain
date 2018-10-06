package main

// Blockchain represents the blockchain structure
type Blockchain struct {
	blocks []*Block
}

// AddBlock fetches the last block's hash and calls NewBlock to create a new block with this info
// The new block gets appended to the Blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewGenesisBlock does exactly what it says
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain creates and returns a new Blockchain type structure with the Genesis Block included
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
