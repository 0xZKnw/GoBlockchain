package main

import (
	"log"
	"net/http"

	"GoBlockchain/common"
)

func main() {

	bc := common.Blockchain {
		Chain: []common.Block{},
	}

	bcS := make(common.BlockchainState)

	gen := common.Block {
		PreviousHash: "0",
		Timestamp: 0,
		Nonce: 0,
		Difficulty: 4,
		Transaction: []common.Transaction{},
	}

	common.Mining(&gen, &bc)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		common.HandleConnection(w, r, &bcS)
	})
	log.Println("websocket node start on localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}