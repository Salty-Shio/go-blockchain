# Proof Of Work

## The lifeblood of block chain

Block chain uses proof of work to generate new blocks. This means that users must do some form of work to recieve some cryptocurrency. This is done by finding the hash for a block. A hash is a one way funciton which creates a fixed length hash from data. Why is this important? This means that a user cannot look at the hash and reverse engineer to get the answer. It should also be noted that because the hash of the previous block is contained in each block, one must not only decifer the current hash, but all previous block hashses to make edits to the network. This is where the security comes from.

## Hashcash

Hashcash is an algorithm that Bitcoin uses for proof of work. The idea is that you get some known data and add a counter so that you have `data + counter`. Next, a criteria for the hash is set forth and your computer must find `data + counter` such that that criteria is met. This algorithm was actually made to stop mas spam emails. They way it did this was by adding a large computation cost to sending an email, thus dettering spammers from sending out mass emails by increasing the time and cost of sending each emial.

In bitcoin, the criteria is adjusted over time  to maintain a constant rate of new blocks (about 10 per minute). The constraint can look like something such as, "Find a hash with the first twenty characters being 0s.

## Proof of Work

The counter during this algorithm is called a *nonce*, or a *number used once*. In our updated code, instead of just hashing the block and adding it to the block chain, it now needs to find a valid hash for the block. This means that the new hash must meet certain criteria. To "mine" the block, the computer adds a nonce to the end of the data and hashes it. If the hash meets the criteria, the nonce is added to the block as well as the valid hash. This allows other computers to hash the data and check against the provided hash trivially, while the task of finding the nonce is non-trivial.