package common

import (
	"encoding/hex"
	"log"

	eciesgo "github.com/ecies/go/v2"
	"github.com/gorilla/websocket"
)

type DAcc struct {
	PrivKey string
	PubKey  string
	Tokens  int
}

// NewDAcc creates a new digital account with automatically generated keys
func NewDAcc(initialTokens int) *DAcc {
	acc := &DAcc{
		Tokens: initialTokens,
	}
	SetKeys(acc)
	return acc
}

func Send(from *DAcc, to *DAcc, amount int, data string, b *Block, bs *BlockchainState) {
	from.Tokens -= amount
	to.Tokens += amount

	// Update blockchain state with new balances
	(*bs)[from.PubKey] = from.Tokens
	(*bs)[to.PubKey] = to.Tokens

	t := Transaction{
		From:   from.PubKey,
		To:     to.PubKey,
		Amount: amount,
		Data:   data,
	}
	b.Transaction = append(b.Transaction, t)
}

func SetKeys(a *DAcc) {
	priv, _ := eciesgo.GenerateKey()
	pub := priv.PublicKey
	a.PrivKey = priv.Hex()
	a.PubKey = hex.EncodeToString(pub.Bytes(true))
}

func WsClient(url string, message string) {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	err = c.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write:", err)
		return
	}

	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("recv: %s", msg)
}
