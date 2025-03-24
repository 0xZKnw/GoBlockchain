package common

import (
    "encoding/hex"
    "log"

    "github.com/ecies/go/v2"
    "github.com/gorilla/websocket"
)

type DAcc struct {
    PrivKey string
    PubKey  string
    Tokens  int
}

func Send(from *DAcc, to *DAcc, amount int, data string, b *Block, bs *BlockchainState) {
    from.Tokens -= amount
    to.Tokens += amount
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