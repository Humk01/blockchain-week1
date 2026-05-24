package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
	"fmt"
)

// Block is one record in the chain.
type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
}

type Transaction struct {
	To string 
	From string
	amount float64
}

func (tx Transaction) String() string {
	return fmt.Sprintf("%s->%s:%.2f", tx.From, tx.To, tx.amount)
}

func (tx Transaction) Hash() string {
	hash := sha256.Sum256([]byte(tx.String()))
	return hex.EncodeToString(hash[:])
}

// CalculateHash returns the SHA-256 hash for the block contents.
func (b Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PreviousHash
	sum := sha256.Sum256([]byte(record))
	return hex.EncodeToString(sum[:])
}

// NewBlock creates a new block that points at the previous hash.
func NewBlock(index int, data, previousHash string) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now().UTC().Format(time.RFC3339Nano),
		Data:         data,
		PreviousHash: previousHash,
	}
	block.Hash = block.CalculateHash()
	return block
}

// NewGenesisBlock creates the first block in a chain.
func NewGenesisBlock() Block {
	return NewBlock(0, "James sent 10 BTC to Gladys", "0")
}
