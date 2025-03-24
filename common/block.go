package common

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
)

type Block struct {
	Hash string
	PreviousHash string
	Timestamp int64
	Nonce int64
	Difficulty int
	Transaction []Transaction
}

type Transaction struct {
	From string
	To string
	Amount int
	Data string
}

func setHash(b *Block) {
	bts, err := json.Marshal(b.Transaction)
	if err != nil {
		panic(err)
	}

	data := []byte{}
	data = append(data, []byte(b.PreviousHash)...)
	data = append(data, []byte(strconv.FormatInt(b.Timestamp, 10))...)
	data = append(data, []byte(strconv.FormatInt(b.Nonce, 10))...)
	data = append(data, []byte(strconv.FormatInt(int64(b.Difficulty), 10))...)
	data = append(data, bts...)

	hash := sha256.Sum256(data)
	b.Hash = hex.EncodeToString(hash[:])
}