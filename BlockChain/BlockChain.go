package main

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) addBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := newBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func newBlockchain() *Blockchain {
	return &Blockchain{[]*Block{newGenesisBlock()}}
}

func newGenesisBlock() *Block {
	return newBlock("Genesis Block", []byte{})
}
