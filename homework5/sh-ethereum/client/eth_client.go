package client

import (
	"encoding/json"
	"math/big"
	"sh-ethereum/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type EthClient struct {
	// 这里可\以添加数据库连接等依赖
	ginContext *gin.Context
	etchClient *ethclient.Client
}

func NewEthClient(c *gin.Context, d *ethclient.Client) *EthClient {
	return &EthClient{ginContext: c, etchClient: d}
}

func (p *EthClient) getBlockByTag() (*types.Header, common.Hash, *utils.AppError) {

	var raw json.RawMessage

	err := p.etchClient.Client().CallContext(p.ginContext, &raw, "eth_getBlockByTag", "latest", false)
	if err != nil {
		return nil, common.Hash{}, utils.NewAppError("EthClient.getBlockByTag", "Failed to call eth_getBlockByTag", err.Error(), 500)
	}

	if len(raw) == 0 || string(raw) == "null" {
		return nil, common.Hash{}, nil
	}

	// 解析完整的区块头字段
	var blockData struct {
		Number      string         `json:"number"`
		Hash        common.Hash    `json:"hash"`
		ParentHash  common.Hash    `json:"parentHash"`
		UncleHash   common.Hash    `json:"sha3Uncles"`
		Coinbase    common.Address `json:"miner"`
		Root        common.Hash    `json:"stateRoot"`
		TxHash      common.Hash    `json:"transactionsRoot"`
		ReceiptHash common.Hash    `json:"receiptsRoot"`
		Bloom       hexutil.Bytes  `json:"logsBloom"`
		Difficulty  *hexutil.Big   `json:"difficulty"`
		GasLimit    hexutil.Uint64 `json:"gasLimit"`
		GasUsed     hexutil.Uint64 `json:"gasUsed"`
		Time        hexutil.Uint64 `json:"timestamp"`
		Extra       hexutil.Bytes  `json:"extraData"`
		MixDigest   common.Hash    `json:"mixHash"`
		Nonce       hexutil.Bytes  `json:"nonce"`
		BaseFee     *hexutil.Big   `json:"baseFeePerGas"`
	}

	if err := json.Unmarshal(raw, &blockData); err != nil {
		return nil, common.Hash{}, utils.NewAppError("EthClient.getBlockByTag", "Failed to unmarshal block data", err.Error(), 500)
	}

	//解析区块号
	num, ok := new(big.Int).SetString(blockData.Number[2:], 16)
	if !ok {
		return nil, common.Hash{}, utils.NewAppError("EthClient.getBlockByTag", "Failed to parse block number", "invalid block number format", 500)
	}

	// 构造完整的 Header
	header := &types.Header{
		ParentHash:  blockData.ParentHash,
		UncleHash:   blockData.UncleHash,
		Coinbase:    blockData.Coinbase,
		Root:        blockData.Root,
		TxHash:      blockData.TxHash,
		ReceiptHash: blockData.ReceiptHash,
		Bloom:       types.BytesToBloom(blockData.Bloom),
		Difficulty:  big.NewInt(0),
		Number:      num,
		GasLimit:    uint64(blockData.GasLimit),
		GasUsed:     uint64(blockData.GasUsed),
		Time:        uint64(blockData.Time),
		Extra:       blockData.Extra,
		MixDigest:   blockData.MixDigest,
		BaseFee:     nil,
	}

	if blockData.Difficulty != nil {
		header.Difficulty = blockData.Difficulty.ToInt()
	}

	if blockData.BaseFee != nil {
		header.BaseFee = blockData.BaseFee.ToInt()
	}

	if len(blockData.Nonce) >= 8 {
		var nonceBytes [8]byte
		copy(nonceBytes[:], blockData.Nonce[:8])
		header.Nonce = types.BlockNonce(nonceBytes)
	}

	// 返回 Header 和 RPC 提供的 hash
	// 注意：手动构造的 Header 计算出的 hash 可能不准确，因为：
	// 1. RPC 返回的某些字段可能格式不完全匹配 go-ethereum 的内部格式
	// 2. Header 的内部缓存字段可能未正确初始化
	// 因此，我们应该直接使用 RPC 返回的 hash，它与浏览器显示的 hash 一致
	return header, blockData.Hash, nil
}

func (p *EthClient) fetchBlockWithRetry(blockNumber *big.Int, maxRetries int) (*types.Block, *utils.AppError) {

}
