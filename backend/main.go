package main

import (
	"fmt"
	"log"
	"net/http"
)

var blockchain *Blockchain

func main() {

	blockchain = NewBlockchain()

	fmt.Println("Blockchain API running on http://localhost:8080")

	http.HandleFunc("/add", AddTransactionHandler)
	http.HandleFunc("/mine", MineBlockHandler)
	http.HandleFunc("/chain", GetBlockchainHandler)
	http.HandleFunc("/search", SearchHandler)
	http.HandleFunc("/pending", GetPendingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
