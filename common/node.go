package common

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleConnection(w http.ResponseWriter, r *http.Request, bcS *BlockchainState, bc *Blockchain) {
	Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	defer ws.Close()

	fmt.Println("Noeud en ligne")

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			panic(err)
		}

		msg := string(p)
		msgS := strings.Split(msg, " ")

		fmt.Println("msg recu : " + msg)
		if msgS[0] == "send" {
			from := NewDAcc(10)
			to := NewDAcc(0)

			amount := 5
			if len(msgS) > 1 {
				if parsedAmount, err := strconv.Atoi(msgS[1]); err == nil {
					amount = parsedAmount
				}
			}

			newBlock := &Block{
				Difficulty: 5,
				Nonce:      0,
				Timestamp:  time.Now().Unix(),
			}

			Send(from, to, amount, "salut", newBlock, bcS)

			Mining(newBlock, bc)
		}
	}
}
