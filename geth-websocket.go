package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
)

type PendingTx struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	Hash             string `json:"hash"`
	To               string `json:"to"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	R                string `json:"r"`
	S                string `json:"s"`
	V                string `json:"v"`
	TransactionIndex string `json:"transactionIndex"`
	Type             string `json:"type"`
	Value            string `json:"value"`
}

// subscribePendings runs in its own goroutine and maintains
// a subscription for new blocks.
func subscribePendings(client *rpc.Client, subch chan PendingTx) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Subscribe to new blocks.
	sub, err := client.EthSubscribe(ctx, subch, "alchemy_pendingTransactions")
	if err != nil {
		fmt.Println("subscribe error:", err)
		return
	}

	// The subscription will deliver events to the channel. Wait for the
	// subscription to end for any reason, then loop around to re-establish
	// the connection.
	fmt.Println("connection lost: ", <-sub.Err())
}

// print hashes of incoming pending transactions
func alchemyPendingStream() {
	url := "wss://eth-mainnet.alchemyapi.io/v2/demo"
	var err error
	subch := make(chan PendingTx)
	client, err := rpc.Dial(url)
	if err != nil {
		fmt.Printf("subscribe err: %e", err)
		return
	}

	// Ensure that subch receives the latest pending.
	go func() {
		for i := 0; ; i++ {
			if i > 0 {
				time.Sleep(2 * time.Second)
			}
			subscribePendings(client, subch)
		}
	}()

	// Print events from the subscription as they arrive.
	for pending := range subch {
		fmt.Println("latest pending:", pending.Hash)
	}
}
