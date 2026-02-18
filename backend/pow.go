package main

import (
	"strings"
)

const Difficulty = 3

func (b *Block) MineBlock() {
	target := strings.Repeat("0", Difficulty)

	for {
		b.Hash = b.CalculateHash()
		if b.Hash[:Difficulty] == target {
			break
		}
		b.Nonce++
	}
}
