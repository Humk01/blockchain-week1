package blockchain

import "testing"

func TestNewChainIsValid(t *testing.T) {
	chain := New()
	chain.AddBlock("Alice sent 10 BTC to Bob")
	chain.AddBlock("Bob sent 5 BTC to Charlie")

	if !chain.IsValid() {
		t.Fatal("expected chain to be valid")
	}
}

func TestTamperingInvalidatesChain(t *testing.T) {
	chain := New()
	chain.AddBlock("Alice sent 10 BTC to Bob")

	chain.Blocks[1].Transactions[0].Data = "Mallory sent 999 BTC to Mallory"

	if chain.IsValid() {
		t.Fatal("expected tampered chain to be invalid")
	}
}
