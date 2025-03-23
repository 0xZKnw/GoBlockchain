package main
import(
	"encoding/hex"

	"github.com/ecies/go/v2"
)

type dAcc struct {
	PrivKey string
	PubKey string
	Tokens int
}

func setKeys(a *dAcc) {
	priv, _ := eciesgo.GenerateKey()
	pub := priv.PublicKey
	a.PrivKey = priv.Hex()
	a.PubKey = hex.EncodeToString(pub.Bytes(true)) 
}

func send(from *dAcc, to *dAcc, amount int, data string, b *Block) {
	from.Tokens -= amount
	to.Tokens += amount
	t := Transaction {
		From: from.PubKey,
		To: to.PubKey,
		Amount: amount,
		Data: data,
	}
	b.Transaction = append(b.Transaction, t)
}