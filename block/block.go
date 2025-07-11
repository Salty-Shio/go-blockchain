package blocks

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

// For simplicity, all of this is combined in this implementation.
type Block struct {
	Timestamp 			int64  // When the block was created
	Data 				[]byte // Data stored in the block
	PreviousBlockHash 	[]byte // Hash of the previous block
	Hash 				[]byte // Hash of the current block
	Nonce 				int    // Nonce for proof of work
}

// Sets the hash of the current block.
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PreviousBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// Generator function to create a new block easily.
func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevHash, []byte{}, 0}
	pow := NewProofOfWork(block) // Create a new proof of work instance
	nonce, hash := pow.Run() // Run the proof of work algorithm

	block.Hash = hash // Set the block's hash
	block.Nonce = nonce

	return block
}


// Special constructor for the genesis block.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}



// PROOF OF WORK

// Difficulty level for the proof of the work algorithms.
const targetBits = 24

// ProofOfWork represents the proof of work algorithm used to mine blocks in the blockchain.
type ProofOfWork struct {
	block  *Block // The block being mined
	target *big.Int // The target value for the proof of work
}

// Our proof of work creates a big int and then shifts it left by the number of bits we want to target.
// Any valid hash must be less than the target value. The smaller the target, the less likely a hash will be valid.
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1) // Initialize target to 1
	// Left Shift the target by (256 - targetBits) bits
	target.Lsh(target, uint(256 - targetBits))

	// Create a new ProofOfWork instance with the block and target
	pow := &ProofOfWork{b, target}

	return pow
}


// This function packages the block data into a byte slice that will be used to calculate the hash.
// Nonce is the number that will be incremented until a valid hash is found.
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PreviousBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Executes the proof of work algorithm.
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0 // Start nonce at 0 and increment it until a valid hash is found

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce) // Get the data to hash
		hash = sha256.Sum256(data)
		// fmt.Printf("\r%x", hash) // Print the hash in hexadecimal format
		hashInt.SetBytes(hash[:]) // Convert the hash to a big integer

		// Check if the hash is valid or not.
		if hashInt.Cmp(pow.target) == -1 {
			break; // The hash is valid, exit the loop	
		} else {
			nonce++ // Increment the nonce and try again
		}
	}
	fmt.Printf("Block mined! Nonce: %d\n\n", nonce)

	return nonce, hash[:]

}

func IntToHex(num int64) []byte {
	return []byte(strconv.FormatInt(num, 16))
}