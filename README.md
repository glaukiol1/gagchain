# gagchain (Gagcoin)

Experimental blockchain built in Golang. It is still under development and features will be updated regulary.

## Features

- Custom TCP server for communication with nodes
- Custom Database built only to store the current chain
- Custom communication messages to communicate with other nodes
- Verifying node authenticity by its address
- Concurrency
- Easy to compile and run

## Using this for your own blockchain project

If you use this as a template, please do not remove any credits, and I expect to be credited in a README file or something similar. You can always run your own node, just setup the variables that are found in the `main.go` file (root directory).

## Usage

WHen you run a node, you can do the following things;

Init blockchain;

```go
blockchain.InitBlockchain()
```

Or: `&blockchain.Blockchain{chain []*blockchain.Block}`

Init Transaction Pool;

```go
blockchain.NewTransactionPool()
```

Add Block to the blockchain (this also mines it if;

- The Miner field is ""
- config.Mining_Node = true
  
```go
(*blockchain.Blockchain).AddBlock(data []*blockchain.Transaction, hash []byte, nonce int, Miner string)
```

Add transaction to the Transaction Pool;

```go
(*blockchain.TransactionPool).AddTransaction(transaction *blockchain.Transaction)
```

New Transaction;

```go
(*blockchain.Blockchain).NewTransactionInstance(from *ecdsa.PublicKey, to string, amount int)
```