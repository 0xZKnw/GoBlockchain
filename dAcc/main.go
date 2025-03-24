package main

import (
	"GoBlockchain/common"
)

func main() {
	common.WsClient("ws://localhost:8080/ws", "send a1 a2 20 salut")
}
