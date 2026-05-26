package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"blockchain-week1/internal/blockchain"
)

// printHelp shows the available CLI commands.
func printHelp() {
	fmt.Println("Commands: add, print, verify, help, quit")
	fmt.Println("  add    - Add a new block")
	fmt.Println("  print  - View the blockchain")
	fmt.Println("  verify - Check if chain is valid")
	fmt.Println("  help   - Show this help message")
	fmt.Println("  quit   - Exit")
}

// readTransaction reads one transaction from the add flow.
func readTransaction(reader *bufio.Reader) (blockchain.Transaction, bool, error) {
	fmt.Print("to: ")
	to, _ := reader.ReadString('\n')
	to = strings.TrimSpace(to)
	if to == "" {
		return blockchain.Transaction{}, false, fmt.Errorf("to cannot be empty")
	}
	if to == "done" {
		return blockchain.Transaction{}, true, nil
	}

	fmt.Print("from: ")
	from, _ := reader.ReadString('\n')
	from = strings.TrimSpace(from)
	if from == "" {
		return blockchain.Transaction{}, false, fmt.Errorf("from cannot be empty")
	}

	fmt.Print("amount: ")
	amountText, _ := reader.ReadString('\n')
	amountText = strings.TrimSpace(amountText)
	amount, err := strconv.ParseFloat(amountText, 64)
	if err != nil {
		return blockchain.Transaction{}, false, fmt.Errorf("amount must be a number")
	}

	return blockchain.Transaction{To: to, From: from, Amount: amount}, false, nil
}

// runAddMode collects transactions and saves them as a new block.
func runAddMode(reader *bufio.Reader, chain *blockchain.BlockChain) {
	var transactions []blockchain.Transaction
	fmt.Println("fill in the data. type done at the to: prompt to return to the main CLI")

	for {
		tx, done, err := readTransaction(reader)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if done {
			if len(transactions) > 0 {
				chain.AddBlock(transactions)
			}
			return
		}
		transactions = append(transactions, tx)
		fmt.Println("Press Enter at the next to: prompt, or type done to finish this block.")
		fmt.Println()
	}
}

// main runs the CLI loop.
func main() {
	chain := blockchain.New()
	chain.AddBlock([]blockchain.Transaction{{From: "Alice", To: "Bob", Amount: 10}})
	chain.AddBlock([]blockchain.Transaction{{From: "Bob", To: "Charlie", Amount: 5}})
	chain.AddBlock([]blockchain.Transaction{{From: "Charlie", To: "Alice", Amount: 2}})

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple Blockchain CLI")
	printHelp()

	for {
		fmt.Print("\nmain> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "add":
			runAddMode(reader, chain)
			fmt.Println("Back to main CLI. Type help to see commands.")
		case "print":
			chain.Print()
		case "verify":
			if chain.IsValid() {
				fmt.Println("Blockchain is VALID")
			} else {
				fmt.Println("Blockchain is INVALID - tampering detected!")
			}
		case "help":
			printHelp()
		case "quit":
			fmt.Println("Goodbye")
			return
		default:
			fmt.Println("Unknown command. Type 'help' to see available commands.")
		}
	}
}
