package common

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func HandleConnection(w http.ResponseWriter, r *http.Request, bcS *BlockchainState) {
	Upgrader.CheckOrigin = func (r *http.Request) bool {return true}
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
			Send(&DAcc{}, &DAcc{}, 5, "salut", &Block{}, bcS)
		}
	}
}