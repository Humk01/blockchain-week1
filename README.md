# Simple Blockchain in Go

A basic blockchain implementation to understand how hash chains work.

## Structure
- `main.go` - CLI entrypoint
- `internal/blockchain` - block and chain logic
- `internal/blockchain/blockchain_test.go` - core behavior tests

## What It Does
- Creates blocks with multiple transactions
- Uses a nested `add` flow for building a block
- Prompts for transaction fields one at a time: `to`, `from`, `amount`
- Builds a Merkle tree for each block
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
- `help` - Show the available commands
- `quit` - Exit

## Transaction Input
When you run `add`, the CLI asks for:
- `To`
- `From`
- `Amount`

Type `done` at the `to:` prompt to finish the block and return to the main CLI.
After that, the CLI returns to the `main>` prompt.

## What I Learned
- Hash functions create unique fingerprints of data
- Changing old data breaks the chain
- This is the foundation of Bitcoin's immutability
