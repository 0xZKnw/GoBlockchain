package main
import (
	"fmt"
)

func main() {

	bc := Blockchain {
		Chain: []Block{},
	}

	a := dAcc {
		Tokens: 0,
	}
	setKeys(&a)

	b := Block {
		PreviousHash: "0",
		Timestamp: 0,
		Nonce: 0,
		Difficulty: 4,
		Transaction: []Transaction{{From: "E", To: "C", Amount: 40, Data: "salut chef!"}},
	}
	Mining(&b)
	bc.Chain = append(bc.Chain, b)
	fmt.Println(bc.Chain[0].Hash)
}