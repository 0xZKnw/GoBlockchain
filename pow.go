package main
import (
	"fmt"
	"strings"
)

func Mining(b *Block) {
	prefix := strings.Repeat("0", b.Difficulty)
	for !strings.HasPrefix(b.Hash, prefix) {
		b.Nonce += 1
		setHash(b)
		fmt.Println(b.Nonce)
	}
}