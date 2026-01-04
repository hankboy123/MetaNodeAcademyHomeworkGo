package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const erc20ABIJSON = `[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "from",
				"type": "address"
			},			
			{
				"indexed": true,
				"name": "to",
				"type": "address"
			},			
			{
				"indexed": false,
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Transfer",
		"type": "event"
	}]`

type TransferEvent struct {
	BlockNumber uint64    `json:"block_number"`
	TxHash      string    `json:"tx_hash"`
	From        string    `json:"from"`
	To          string    `json:"to"`
	Value       string    `json:"value"`
	Timstamp    time.Time `json:"timestamp"`
}

type EventStore struct {
	mu     sync.RWMutex
	Events []TransferEvent
	limit  int
}

func NewEventStore(limit int) *EventStore {
	return &EventStore{
		Events: make([]TransferEvent, 0, limit),
		limit:  limit,
	}
}

func (es *EventStore) Add(event TransferEvent) {
	es.mu.Lock()
	defer es.mu.Unlock()

	if len(es.Events) >= es.limit {
		es.Events = es.Events[1:]
	}
	es.Events = append(es.Events, event)
}

func (es *EventStore) List() []TransferEvent {
	es.mu.RLock()
	defer es.mu.RUnlock()

	eventsCopy := make([]TransferEvent, len(es.Events))
	copy(eventsCopy, es.Events)
	return eventsCopy
}

func main() {
	os.Setenv("ETH_HTTP_URL", "https://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY")
	os.Setenv("ETH_WS_URL", "wss://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY")
	os.Setenv("ERC20_CONTRACT_ADDRESS", "wss://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY")

	rpcURL := os.Getenv("ETH_WS_URL")

	if rpcURL == "" {
		rpcURL = os.Getenv("ETH_HTTP_URL")
	}

	if rpcURL == "" {
		log.Fatal("ETH_WS_URL environment variable is not set")
	}

	contractHex := os.Getenv("ERC20_CONTRACT_ADDRESS")
	if contractHex == "" {
		log.Fatal("ERC20_CONTRACT_ADDRESS environment variable is not set")
	}
	contractAddress := common.HexToAddress(contractHex)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	parsedABI, err := abi.JSON(strings.NewReader(erc20ABIJSON))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	store := NewEventStore(100)

	go subscribeTransferEvents(ctx, client, contractAddress, parsedABI, store)

	mux := http.NewServeMux()
	mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		events := store.List()
		_ = json.NewEncoder(w).Encode(events)
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		log.Printf("HTTP server listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server error: %v", err)
		}
	}()

	// 优雅退出
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh
	fmt.Printf("received signal %s, shutting down...\n", sig.String())

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	_ = server.Shutdown(shutdownCtx)
	cancel()
}

func subscribeTransferEvents(ctx context.Context, client *ethclient.Client, contract common.Address, parsedABI abi.ABI, store *EventStore) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contract},
	}
	logsCh := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(ctx, query, logsCh)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}

	log.Printf("listening Transfer events of %s", contract.Hex())
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case vLog := <-logsCh:
			if len(vLog.Topics) == 0 {
				continue
			}

			var transferEvent struct {
				From  common.Address
				To    common.Address
				Value *big.Int
			}
			err := parsedABI.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Printf("Failed to unpack log: %v", err)
				continue
			}

			if len(vLog.Topics) >= 3 {
				transferEvent.From = common.BytesToAddress(vLog.Topics[1].Bytes())
				transferEvent.To = common.BytesToAddress(vLog.Topics[2].Bytes())
			}

			event := TransferEvent{
				BlockNumber: vLog.BlockNumber,
				TxHash:      vLog.TxHash.Hex(),
				From:        transferEvent.From.Hex(),
				To:          transferEvent.To.Hex(),
				Value:       transferEvent.Value.String(),
				Timstamp:    time.Now(),
			}

			store.Add(event)
			log.Printf("New Transfer event: %+v", event)
		case <-ctx.Done():
			log.Println("context cancelled, stop subscription")
			return
		}
	}

}
