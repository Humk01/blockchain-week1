package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// Block is one record in the chain.
type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	MerkleRoot   string
	PreviousHash string
	Hash         string
}

// Transaction is a simple payload that participates in the Merkle tree.
type Transaction struct {
	To     string
	From   string
	Amount float64
}

// String formats a transaction for display and hashing.
func (tx Transaction) String() string {
	return fmt.Sprintf("%s sent %.2f to %s", tx.From, tx.Amount, tx.To)
}

// Hash returns the SHA-256 hash of a transaction.
func (tx Transaction) Hash() string {
	sum := sha256.Sum256([]byte(tx.String()))
	return hex.EncodeToString(sum[:])
}

// CalculateHash returns the SHA-256 hash for the block contents.
// CalculateHash returns the SHA-256 hash for the block contents.
func (b Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + calculateMerkleRoot(b.Transactions) + b.PreviousHash
	sum := sha256.Sum256([]byte(record))
	return hex.EncodeToString(sum[:])
}

// calculateMerkleRoot builds a Merkle root from the provided transactions.
// calculateMerkleRoot combines transaction hashes into one Merkle root.
func calculateMerkleRoot(transactions []Transaction) string {
	if len(transactions) == 0 {
		return ""
	}

	hashes := make([]string, 0, len(transactions))
	for _, tx := range transactions {
		hashes = append(hashes, tx.Hash())
	}

	for len(hashes) > 1 {
		var nextLevel []string

		for i := 0; i < len(hashes); i += 2 {
			combined := hashes[i]
			if i+1 < len(hashes) {
				combined += hashes[i+1]
			} else {
				combined += hashes[i]
			}

			sum := sha256.Sum256([]byte(combined))
			nextLevel = append(nextLevel, hex.EncodeToString(sum[:]))
		}

		hashes = nextLevel
	}

	return hashes[0]
}

// NewBlock creates a new block that points at the previous hash.
// NewBlock creates a block that links to the previous hash.
func NewBlock(index int, previousHash string, transactions []Transaction) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now().UTC().Format(time.RFC3339Nano),
		Transactions: transactions,
		MerkleRoot:   calculateMerkleRoot(transactions),
		PreviousHash: previousHash,
	}
	block.Hash = block.CalculateHash()
	return block
}

// NewGenesisBlock creates the first block in a chain.
// NewGenesisBlock creates the first block in the chain.
func NewGenesisBlock() Block {
	return NewBlock(0, "0", []Transaction{{From: "James", To: "Gladys", Amount: 10}})
}
