package client

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthListener struct {
	context   *context.Context
	rpcClient *ethclient.Client
}

func (e *EthListener) Listen() {

	headers := make(chan *types.Header)
	sub, err := e.rpcClient.SubscribeNewHead(*e.context, headers)
	if err != nil {
		println("Failed to subscribe to new headers:", err.Error())
		return
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case h := <-headers:
			if h == nil {
				continue
			}
			fmt.Printf("[%s] New Block - Number: %d, Hash: %s\n",
				time.Now().Format(time.RFC3339),
				h.Number.Uint64(),
				h.Hash().Hex(),
			)
		case err := <-sub.Err():
			println("Subscription error:", err.Error())
			return
		case sig := <-sigCh:
			// Handle graceful shutdown
			fmt.Printf("received signal %s, shutting down...\n", sig.String())
			return
		case <-(*e.context).Done():
			fmt.Println("context cancelled, shutting down...")
			return

		}
	}
}
