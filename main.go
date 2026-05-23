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


func (b *Block) CalculateHash() string {

	// combines all blocks data into one string 
	record := strconv.Itoa(b.INDEX) + b.Timestamp + b.Data + b.PreviousHash + b.Hash
	hash := sha256.Sum256([]byte(record))

	//convert to hex string
	return hex.EncodeToString(hash[:])


}

func createGenesisBlock() Block {
	genesisBlock := Block {
		INDEX: 0,
		Timestamp: time.Now().String(),
		Data: "O happy day",
		PreviousHash: "0",
	}
	genesisBlock.Hash = genesisBlock.CalculateHash()
	return genesisBlock
}

func main() {
	genesis := createGenesisBlock()
    fmt.Printf("Block #%d created\n", genesis.INDEX)
    fmt.Printf("Data: %s\n", genesis.Data)
    fmt.Printf("Hash: %s\n", genesis.Hash)
    fmt.Printf("Previous Hash\n: %s", genesis.PreviousHash)
}