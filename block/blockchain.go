package block

// For now we store the blocks in a slice so we have an ordered list of blocks.
type BlockChain struct {
	// Blocks should be private (blocks) but is public for testing.
	Blocks []*Block // Slice of pointers pointing to blocks.
}

// Adds blocks to the blockchain.
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1] // Get the last block on the chain.
	newBlock := NewBlock(data, prevBlock.Hash) // Create the new block.
	bc.Blocks = append(bc.Blocks,newBlock) // Append the new block to the chain.
}

// Constructs a new blockchain.
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}