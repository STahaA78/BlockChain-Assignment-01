package main

import "fmt"

type Blockchain struct {
	Blocks              []*Block
	PendingTransactions []string
}

// Create new blockchain with Genesis block
func NewBlockchain() *Blockchain {

	genesis := NewBlock(0, []string{"22L-6638"}, "0") // Roll number required
	genesis.MineBlock()

	return &Blockchain{
		Blocks:              []*Block{genesis},
		PendingTransactions: []string{},
	}
}

// Add transaction to pending pool
func (bc *Blockchain) AddTransaction(data string) {
	bc.PendingTransactions = append(bc.PendingTransactions, data)
}

// Mine pending transactions into a new block
func (bc *Blockchain) MinePendingTransactions() {

	if len(bc.PendingTransactions) == 0 {
		fmt.Println("No transactions to mine")
		return
	}

	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := NewBlock(
		prevBlock.Index+1,
		bc.PendingTransactions,
		prevBlock.Hash,
	)

	newBlock.MineBlock()

	bc.Blocks = append(bc.Blocks, newBlock)

	// Clear pending transactions
	bc.PendingTransactions = []string{}
}
