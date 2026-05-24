# Simple Blockchain in Go

A basic blockchain implementation to understand how hash chains work.

## Structure
- `main.go` - CLI entrypoint
- `internal/blockchain` - block and chain logic
- `internal/blockchain/blockchain_test.go` - core behavior tests

## What It Does
- Creates blocks with data
- Links blocks using cryptographic hashes
- Detects tampering with old blocks

## How to Run
```bash
go run .
```

## Commands
- `add` - Add a new block
- `print` - View the blockchain
- `verify` - Check if chain is valid
- `quit` - Exit

## What I Learned
- Hash functions create unique fingerprints of data
- Changing old data breaks the chain
- This is the foundation of Bitcoin's immutability
