package blockchain

import "testing"

// TestNewChainIsValid checks that a fresh chain stays valid after adding blocks.
func TestNewChainIsValid(t *testing.T) {
	chain := New()
	chain.AddBlock([]Transaction{{From: "Alice", To: "Bob", Amount: 10}})
	chain.AddBlock([]Transaction{{From: "Bob", To: "Charlie", Amount: 5}})

	if !chain.IsValid() {
		t.Fatal("expected chain to be valid")
	}
}

// TestTamperingInvalidatesChain checks that editing a transaction breaks the chain.
func TestTamperingInvalidatesChain(t *testing.T) {
	chain := New()
	chain.AddBlock([]Transaction{{From: "Alice", To: "Bob", Amount: 10}})

	chain.Blocks[1].Transactions[0].Amount = 999

	if chain.IsValid() {
		t.Fatal("expected tampered chain to be invalid")
	}
}
