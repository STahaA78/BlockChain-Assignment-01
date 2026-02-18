package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index      int
	Timestamp  string
	Data       []string
	PrevHash   string
	Hash       string
	Nonce      int
	MerkleRoot string
}

// Create new block
func NewBlock(index int, data []string, prevHash string) *Block {
	block := &Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
		Nonce:     0,
	}

	block.MerkleRoot = GenerateMerkleRoot(data)
	return block
}

// Calculate hash of block
func (b *Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) +
		b.Timestamp +
		b.PrevHash +
		b.MerkleRoot +
		strconv.Itoa(b.Nonce)

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}
