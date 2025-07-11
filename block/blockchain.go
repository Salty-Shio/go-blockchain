package blocks

type BlockChain struct {
	Blocks []*Block // Slice of pointers pointing to Blocks.
}

// Adds Blocks to the blockchain.
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1] // Get the last block on the chain.
	newBlock := NewBlock(data, prevBlock.Hash) // Create the new block.
	bc.Blocks = append(bc.Blocks,newBlock) // Append the new block to the chain.
}

// Constructs a new blockchain.
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}