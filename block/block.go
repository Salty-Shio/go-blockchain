package blocks

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
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

// Serializes the block into a byte slice using gob encoding.
// This allows the block to be easily stored in the database.
func (b * Block) Serialize() []byte {
	var result bytes.Buffer // Create a new buffer to hold the serialized data
	encoder := gob.NewEncoder(&result) // Create a new gob encoder
	err := encoder.Encode(b) // Encode the block into the buffer

	// Error handling
	if err != nil {
		fmt.Println("Error encoding block:", err)
		return nil // Return nil if there was an error
	}

	return result.Bytes()
}

// HELPERS

// Generator function to create a new block easily.
func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevHash, []byte{}, 0}
	pow := NewProofOfWork(block) 
	nonce, hash := pow.Run() 

	block.Hash = hash 
	block.Nonce = nonce

	return block
}


// Special constructor for the genesis block.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func DeserializeBlock(serializedBlock []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(serializedBlock))
	err := decoder.Decode(&block) // Decode the serialized block into a Block struct
	if err != nil {
		fmt.Println("Error decoding block:", err)
		return nil // Return nil if there was an error
	}

	return &block // Return a pointer to the decoded block
}

// PROOF OF WORK

const targetBits = 24

// ProofOfWork represents the proof of work algorithm used to mine blocks in the blockchain.
type ProofOfWork struct {
	block  *Block // The block being mined
	target *big.Int // The target value for the proof of work
}

// Our proof of work creates a big int and then shifts it left by the number of bits we want to target.
// Any valid hash must be less than the target value. The smaller the target, the less likely a hash will be valid.
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1) 
	target.Lsh(target, uint(256 - targetBits)) 
	pow := &ProofOfWork{b, target} 
	return pow
}


// Prepares data to be hashed for the proof of work algorithm.
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
		data := pow.prepareData(nonce) 
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:]) 

		if hashInt.Cmp(pow.target) == -1 {
			break; 
		} else {
			nonce++ 
		}
	}
	fmt.Printf("Block mined! Nonce: %d\n\n", nonce)

	return nonce, hash[:]
}

func IntToHex(num int64) []byte {
	return []byte(strconv.FormatInt(num, 16))
}