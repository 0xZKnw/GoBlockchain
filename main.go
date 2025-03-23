package main
import (
	"fmt"
)

func main() {

	bc := Blockchain {
		Chain: []Block{},
	}

	a := dAcc {
		Tokens: 10,
	}
	setKeys(&a)

	a2 := dAcc {
		Tokens: 0,
	}
	setKeys(&a2)

	b := Block {
		PreviousHash: "0",
		Timestamp: 0,
		Nonce: 0,
		Difficulty: 4,
		Transaction: []Transaction{{From: "E", To: "C", Amount: 40, Data: "salut chef!"}},
	}

	send(&a, &a2, 2, "salut", &b)

	Mining(&b)
	bc.Chain = append(bc.Chain, b)
	fmt.Println(bc.Chain[0].Hash)
	fmt.Println(a.Tokens)
	fmt.Println(a2.Tokens)
}