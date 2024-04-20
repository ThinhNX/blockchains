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
	fmt.Println("=====================")
	fmt.Println("*****INFORMATION*****")
	fmt.Println("=====================")

	//after add blocks to chain, now we check the block information
	for _, currentBlock := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", currentBlock.PrevBlockHash)
		fmt.Printf("Data: %s\n", currentBlock.Data)
		fmt.Printf("Hash: %x\n", currentBlock.Hash)
		pow := blocks.NewProofOfWork(currentBlock)
		fmt.Printf("Nonce: %x\n", currentBlock.Nonce)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
	}
}