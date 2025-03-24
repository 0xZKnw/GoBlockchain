package common
import (
	"fmt"
	"strings"
	"strconv"
)

func Mining(b *Block, bc *Blockchain) {
	prefix := strings.Repeat("0", b.Difficulty)
	for !strings.HasPrefix(b.Hash, prefix) {
		setHash(b)
		b.Nonce += 1
	}
	bc.Chain = append(bc.Chain, *b)
	fmt.Println("block n°" + strconv.FormatInt(int64(len(bc.Chain)-1), 10) + " mined")
}