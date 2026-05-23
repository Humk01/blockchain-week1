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

func main() {
	// create blockchain starting with the genesis block
	blockchain := BlockChain{
		Blocks: []Block{createGenesisBlock()},
	}

	// add some blocks
	blockchain.addBlock("Alice sent 10 BTC to Bob")
	blockchain.addBlock("Bob sent 5 BTC to Charlie")
	blockchain.addBlock("Charlie sent 2 BTC to Alice")

	for _, block := range blockchain.Blocks {
		fmt.Printf("Block #%d\n", block.INDEX)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Previous Hash: %s\n\n", block.PreviousHash)
	}
}
