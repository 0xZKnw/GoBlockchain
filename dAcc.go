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