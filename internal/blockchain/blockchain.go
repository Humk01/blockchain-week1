package blockchain

import "fmt"

// BlockChain stores the ordered list of blocks.
type BlockChain struct {
	Blocks []Block
}

// New creates a blockchain with a genesis block.
func New() *BlockChain {
	return &BlockChain{
		Blocks: []Block{NewGenesisBlock()},
	}
}

// AddBlock appends a new block to the chain.
func (bc *BlockChain) AddBlock(data string) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(previousBlock.Index+1, data, previousBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// IsValid checks that every block hash and link is intact.
func (bc *BlockChain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			fmt.Printf("Block #%d has invalid hash!\n", currentBlock.Index)
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			fmt.Printf("Block #%d is not properly linked to previous block!\n", currentBlock.Index)
			return false
		}
	}

	return true
}

// Print writes the chain to stdout in a readable format.
func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("*********************\n")
		fmt.Printf("Block #%d\n", block.Index)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Merkle: %s\n", block.MerkleRoot)

		hash := block.Hash
		if len(hash) > 16 {
			hash = hash[:16] + "..."
		}
		fmt.Printf("Hash: %s\n", hash)

		previousHash := block.PreviousHash
		if len(previousHash) > 16 {
			previousHash = previousHash[:16] + "..."
		}
		fmt.Printf("Previous: %s\n", previousHash)
	}
}
