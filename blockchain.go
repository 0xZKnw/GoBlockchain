package main

type Blockchain struct {
	Chain []Block
}

func (b Blockchain) getBlock(n int) Block {
	return b.Chain[n]
}