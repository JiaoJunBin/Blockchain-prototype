package cli

import "blockchain/core"

type server struct {
	Blockchain *core.Blockchain
	mempool    map[hash]*core.Transaction
}
