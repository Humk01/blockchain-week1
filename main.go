package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	INDEX        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
}

type BlockChain struct {
	Blocks []Block
}

func (b *Block) CalculateHash() string {

	// combines all blocks data into one string
	record := strconv.Itoa(b.INDEX) + b.Timestamp + b.Data + b.PreviousHash + b.Hash
	hash := sha256.Sum256([]byte(record))

	//convert to hex string
	return hex.EncodeToString(hash[:])

}

func createGenesisBlock() Block {
	genesisBlock := Block{
		INDEX:        0,
		Timestamp:    time.Now().String(),
		Data:         "James sent 10 BTC to Gladys",
		PreviousHash: "0",
	}
	genesisBlock.Hash = genesisBlock.CalculateHash()
	return genesisBlock
}

// add new block to make it a blockchain duh
func (bc *BlockChain) addBlock(data string) {
	// last book from the chain
	previousBlock := bc.Blocks[len(bc.Blocks)-1]

	// create new block
	newBlock := Block{
		INDEX:        previousBlock.INDEX + 1,
		Timestamp:    time.Now().String(),
		Data:         data,
		PreviousHash: previousBlock.Hash,
	}
	newBlock.Hash = newBlock.CalculateHash()

	// add to the chain
	bc.Blocks = append(bc.Blocks, newBlock)
}

// this function checks if the blockchain was tampered with

func (bc *BlockChain) isValid() bool {
	// checks each block (starting from block 1)
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			fmt.Printf("Block #%d has invalid hash!\n", currentBlock.INDEX)
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			fmt.Printf("Block #%d is not properly linked to previous block!\n", currentBlock.INDEX)
			return false
		}
	}
	return true
}

func main() {
	// create blockchain starting with the genesis block
	blockchain := BlockChain{
		Blocks: []Block{createGenesisBlock()},
	}

	// add some blocks
	blockchain.addBlock("Alice sent 10 BTC to Bob")
	blockchain.addBlock("Bob sent 5 BTC to Charlie")
	blockchain.addBlock("Charlie sent 2 BTC to Alice")

	// Validate
	fmt.Println("Is blockchain valid?", blockchain.isValid())

	for _, block := range blockchain.Blocks {
		fmt.Printf("Block #%d\n", block.INDEX)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Previous Hash: %s\n\n", block.PreviousHash)
	}

	
}
