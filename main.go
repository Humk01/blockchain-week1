package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func (bc *BlockChain) printChain() {
	for _, block := range bc.Blocks {
		fmt.Printf("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
		fmt.Printf("Block #%d\n", block.INDEX)
		fmt.Printf("Data: %s\n", block.Data)

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

func main() {
	// create blockchain starting with the genesis block
	blockchain := BlockChain{
		Blocks: []Block{createGenesisBlock()},
	}

	// add some blocks
	blockchain.addBlock("Alice sent 10 BTC to Bob")
	blockchain.addBlock("Bob sent 5 BTC to Charlie")
	blockchain.addBlock("Charlie sent 2 BTC to Alice")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple Blockchain CLI")
	fmt.Println("Commands: add, print, verify, quit")

	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "add":
			fmt.Print("Enter block data: ")
			data, _ := reader.ReadString('\n')
			data = strings.TrimSpace(data)
			blockchain.addBlock(data)
		case "print":
			blockchain.printChain()

		case "verify":
			if blockchain.isValid() {
				fmt.Println("Blockchain is VALID")
			} else {
				fmt.Println("Blockchain is INVALID - tampering detected!")
			}
		case "quit":
			fmt.Println("Goodbye")
			return

		default:
			fmt.Println("Unknown command. Try add, print, verify, quit")
		}
	}
}
