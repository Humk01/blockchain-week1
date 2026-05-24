package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"blockchain-week1/internal/blockchain"
)

func main() {
	chain := blockchain.New()
	chain.AddBlock("Alice sent 10 BTC to Bob")
	chain.AddBlock("Bob sent 5 BTC to Charlie")
	chain.AddBlock("Charlie sent 2 BTC to Alice")

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
			chain.AddBlock(data)
		case "print":
			chain.Print()
		case "verify":
			if chain.IsValid() {
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
