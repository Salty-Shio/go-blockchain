package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block repressents a single block in the blockchain.
// In the actual bitcoin specifications, the timestamp, hash, and previous hash
// are all block headers and form a separate data structure. After this,
// transactions (in our case, the data field) is also a separate structure.
// For simplicity, all of this is combined in this implementation.
type Block struct {
	Timestamp 			int64  // When the block was created
	Data 				[]byte // Data stored in the block
	PreviousBlockHash 	[]byte // Hash of the previous block
	Hash 				[]byte // Hash of the current block
}

// In our simplified implentation, the hash will be the SHA-256 hash of the
// concatentation of the block fields.
func (b *Block) SetHash() {
	// Convert the timestamp to a byte slice of base 10 integers
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	// Concatenate the previous hash, data, and timestamp to create the header
	headers := bytes.Join([][]byte{b.PreviousBlockHash, b.Data, timestamp}, []byte{})

	// Calculate the SHA-256 hash of the headers
	hash := sha256.Sum256(headers)

	// Set the block's hash to the calculated hash. [:] converts an array to a slice.
	b.Hash = hash[:]
}

// Generator function to create a new block easily.
func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevHash, []byte{}}
	block.SetHash() // Set the hash of the new block
	return block
}

// Create the first block in a blockchain, known as the genesis block.
// A blockcahin MUST have at least on block to funciton, thus it is necessary
// to instantiate with a genesis block.
// The genesis block has no previous block, so its previous hash is an empty byte slice.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}