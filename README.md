# Blockchain-prototype
a simplified bitcoin-like blockchain system
consists of 1.blockchain_prototype 2.mining 3.transactions 4.mint_coins 5.network

1. Blockchain Prototype: construct the blockchain system according to the 
following structure. 
a) Index: the height of current block. 
b) Data: any data that is included in the block 
c) Timestamp: the creation time of block (seconds from Unix Epoch). 
d) Previous Block Hash: SHA-256 hash of previous block. 
e) Current Block Hash: SHA-256 hash of current block. 

2. Mining: implement a Proof-of-Work algorithm. 
a) Combine all the information in the block and start nonce from 0. 
b) Calculate the SHA-256 hash value of all the information. 
c) If the output is under the target, add the new block to the blockchain. 
d) Otherwise, increment nonce by 1 and repeat step c). 

3. Transaction: 
a) Structure: one transaction consists of a transaction ID, an input, and an 
output. 
b) Transaction ID: the transaction ID is calculated by taking a hash of the 
transaction contents. 
c) Output: the output consists of an address and an amount of coins. 
d) Input: the input consists where the coins are coming from (i.e., previous 
transaction ID and index) along with a signature. 

4. Mint Coins:
a) Coinbase Transaction: each block should contain a coinbase transaction at 
the very first place for transactions to mint 50 coins. 
b) Transaction ID: the transaction ID of coinbase transaction is calculated by 
taking a hash of the transaction contents. 
c) Input: the input of coinbase transaction is zero. 
d) Output: the output consists of an address and the amount of minted coins. 

5. Network: two basic interactions should be realized. 
a) getblock: it is used to get the blocks from the other nodes. 
b) inv: it is used to inform the other nodes what blocks or transactions it has. 