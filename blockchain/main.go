package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	timestamp    time.Time
	transaction  []string
	previousHash []byte
	hash         []byte
}

func main() {
	token_x := []string{"A sent 50 token to B"}
	token_y := Blocks(token_x, []byte{})
	fmt.Println("This is our first genenish block")
	fmt.Println(token_y)
}

func Blocks(transaction []string, previousHash []byte) *Block {
	currentTime := time.Now()
	return &Block{
		timestamp:    currentTime,
		transaction:  transaction,
		previousHash: previousHash,
		hash:         generateNewHash(currentTime, transaction, previousHash),
	}
}

func generateNewHash(time time.Time, transaction []string, previousHash []byte) []byte {
	input := append(previousHash, time.String()...)
	for transaction := range transaction {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]

}

func printBlockchain(block *Block) {
	fmt.Printf("\ttime : %s\n", block.timestamp.String())
	fmt.Printf("\tpreviousHash : %x\n", block.previousHash)
	fmt.Printf("\thash : %x\n", block.hash)
	finalTransaction(block)
}

func finalTransaction(block *Block) {
	fmt.Println("T\tTransaction :")
	for i, transaction := range block.transaction {
		fmt.Printf("\t\t%v: %q\n", i, transaction)
	}
}
