package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block is one record in the chain.
type Block struct {
	Index        int
	Timestamp    string
	Data         string
	Transactions []Transaction
	MerkleRoot   string
	PreviousHash string
	Hash         string
}

// Transaction is a simple payload that participates in the Merkle tree.
type Transaction struct {
	Data string
}

func (tx Transaction) String() string {
	return tx.Data
}

func (tx Transaction) Hash() string {
	sum := sha256.Sum256([]byte(tx.String()))
	return hex.EncodeToString(sum[:])
}

// CalculateHash returns the SHA-256 hash for the block contents.
func (b Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.MerkleRoot + b.PreviousHash
	sum := sha256.Sum256([]byte(record))
	return hex.EncodeToString(sum[:])
}

// calculateMerkleRoot builds a Merkle root from the provided transactions.
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
func NewBlock(index int, data, previousHash string) Block {
	transactions := []Transaction{{Data: data}}
	block := Block{
		Index:        index,
		Timestamp:    time.Now().UTC().Format(time.RFC3339Nano),
		Data:         data,
		Transactions: transactions,
		MerkleRoot:   calculateMerkleRoot(transactions),
		PreviousHash: previousHash,
	}
	block.Hash = block.CalculateHash()
	return block
}

// NewGenesisBlock creates the first block in a chain.
func NewGenesisBlock() Block {
	return NewBlock(0, "James sent 10 BTC to Gladys", "0")
}
