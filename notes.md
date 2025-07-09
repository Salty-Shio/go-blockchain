# Part 1

## Blockchain Basics

The idea of a block chain is a series of connected blocks (wow, who would have thought?). Each of these blocks contains a hash of the previous block, information about the current block, and some form of data being stored,for example, bitcoin stores transational data in it's blocks.

## Blocks

### Our data

Bitcoin, and other forms of cryptocurrency are generally made up of multiple different types of data structures. We will simplify the structure. We will include the timesstamp, data, previous block hash, and the current hash.

### Block Functions

Now that we hav eour data type, we'll put together a few funcitons. First we need a funciton that gets the hash of the current block. This will be used to create a new block. That brings us to the second function, a constructor for a new block. Finally, a blockchain has the requirement of having at least one block, named the genesis block, to function. This is because each block requires a previous block, but there is no block previous to the genesis block. The final function here will just return an empty block with no previous block.

## Block Chain

### Block Chain Type

Now we can create a block chain. For simplicity we'll just make the Block Chain a slice of blocks.

### Methods

First we need a constructor. This  constructor will call the `NewGenesisBlock()` function, and then it will add it to the slice of blocks. Other than that, we need to add to the block chain so we'll add a function that takes in data for a new block and retrieves the previous block hash, then adds a new block to the block chain.

## Conclusion

This is the bare minimum proof of concept for block chain. The main function is a test fuction. To be able to test the code so far, in the BlockCahin struct, blocks were made public, thought they should be private.