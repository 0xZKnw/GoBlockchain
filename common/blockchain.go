package common

type Blockchain struct {
	Chain []Block
}

type BlockchainState map[string]int

func (b Blockchain) getBlock(n int) Block {
	return b.Chain[n]
}