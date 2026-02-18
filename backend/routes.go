package main

import (
	"encoding/json"
	"net/http"
)

// Enable CORS and handle preflight headers
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
}

// Add single transaction to pending pool
func AddTransactionHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	// Handle CORS preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Transaction string `json:"transaction"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	blockchain.AddTransaction(data.Transaction)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Transaction added to pending pool",
	})
}

// Mine pending transactions
func MineBlockHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	// Handle CORS preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	blockchain.MinePendingTransactions()

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Block mined successfully",
	})
}

// View blockchain
func GetBlockchainHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	json.NewEncoder(w).Encode(blockchain.Blocks)
}

// Search transaction
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	query := r.URL.Query().Get("data")

	for _, block := range blockchain.Blocks {
		for _, tx := range block.Data {
			if tx == query {
				json.NewEncoder(w).Encode(block)
				return
			}
		}
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Transaction not found",
	})
}

// View pending transactions
func GetPendingHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	json.NewEncoder(w).Encode(blockchain.PendingTransactions)
}
