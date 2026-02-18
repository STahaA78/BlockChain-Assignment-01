package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateMerkleRoot(transactions []string) string {
	if len(transactions) == 0 {
		return ""
	}

	var hashes []string

	for _, tx := range transactions {
		hash := sha256.Sum256([]byte(tx))
		hashes = append(hashes, hex.EncodeToString(hash[:]))
	}

	for len(hashes) > 1 {
		var newLevel []string

		for i := 0; i < len(hashes); i += 2 {
			if i+1 < len(hashes) {
				combined := hashes[i] + hashes[i+1]
				hash := sha256.Sum256([]byte(combined))
				newLevel = append(newLevel, hex.EncodeToString(hash[:]))
			} else {
				newLevel = append(newLevel, hashes[i])
			}
		}
		hashes = newLevel
	}

	return hashes[0]
}
