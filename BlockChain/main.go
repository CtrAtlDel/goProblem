package main

import "fmt"

func main() {
	bc := newBlockchain()

	bc.addBlock("Send 1 BTC to Ivan")
	bc.addBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlock)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
