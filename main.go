package main

import (
	"fmt"

	"github.com/Salty-Shio/go-blockchain/block"
)

func main() {
	bc := block.NewBlockChain() // Create a new blockchain
	
	// Add a few blocks to the blockchain
	bc.AddBlock("If humans can only hold their breath for 10 minutes then how was Louis Armstrong able to become the first man on the moon?")
	bc.AddBlock("I am in space, on the moon. HELP. I am losing oxygen. I am going to suffocate.")

	// Iterate through blocks in the blockchain
	for _, block := range bc.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PreviousBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Println();
	}
}