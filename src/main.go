package main

import (
	"blockchains/src/blocks"
	"fmt"
	"strconv"
)
func main() {
	bc := blocks.NewBlockChain()

	bc.AddBlock("Send 1 BTC to ThinhNX")
	bc.AddBlock("Send 2 more BTC to ThinhNX")
	fmt.Println("==============================================================")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blocks.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
	}
}