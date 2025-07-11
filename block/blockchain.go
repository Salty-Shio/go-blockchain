package blocks

import (
	"github.com/boltdb/bolt"
)

const blocksBucket = "blocks"

type BlockChain struct {
	tip []byte        // The hash of the latest block
	db  *bolt.DB      // The database where the blockchain is stored
}

// Adds Blocks to the blockchain.
func (bc *BlockChain) AddBlock(data string) {
	// Updated logic to add to database instead of in-memory.
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l")) // Get the last block's hash

		return nil
	})

	if err != nil {
		// Handle the error, possibly by logging it or returning an error.
		panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		if err != nil {
			// Handle the error, possibly by logging it or returning an error.
			panic(err)
		}

		return nil
	})
}

// Constructs a new blockchain.
func NewBlockChain() *BlockChain {
	// Updated logic checks for an existing blockchain first.
	// If it exists, it returns the existing blockchain, otherwise it creates a new one.
	var tip []byte
	db, err := bolt.Open("my.db", 0600, nil)

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				// Handle the error, possibly by logging it or returning an error.
				panic(err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash) // Store the latest block hash
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l")) // Get the latest block hash
		}

		return nil
	})

	if err != nil {
		// Handle the error, possibly by logging it or returning an error.
		panic(err)
	}

	bc := &BlockChain{tip, db}
	return bc
}

// ITERATOR

type BlockChainIterator struct {
	currenthash []byte // The current block's hash
	db          *bolt.DB // The database where the blockchain is stored
}

// Point the iterator to the first block in the blockchain.
func (bc *BlockChain) Iterator() *BlockChainIterator {
	bci := &BlockChainIterator{bc.tip, bc.db}

	return bci
}

// Get the next block in the blockchain.
func (i *BlockChainIterator) Next() *Block {
	var block *Block
	
	err := i.db.View(func(tx *bolt.Tx) error {
		b:= tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currenthash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		panic(err)
	}

	i.currenthash = block.PreviousBlockHash

	return block
}
