package suchain

import (
	"context"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
	"math/big"
)

type WalletManager struct {
	openwallet.AssetsAdapterBase

	Config          *WalletConfig                   // 节点配置
	Decoder         openwallet.AddressDecoder       //地址编码器
	TxDecoder       openwallet.TransactionDecoder   //交易单编码器
	Log             *log.OWLogger                   //日志工具
	ContractDecoder openwallet.SmartContractDecoder //智能合约解析器
	Blockscanner    *SGUBlockScanner                //区块扫描器
	Api             *Api                            //本地封装的http client
	Context         context.Context
}

func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.Config = NewConfig(Symbol)
	wm.Blockscanner = NewSGUBlockScanner(&wm)
	wm.Decoder = NewAddressDecoder(&wm)
	wm.TxDecoder = NewTransactionDecoder(&wm)
	wm.Log = log.NewOWLogger(wm.Symbol())

	wm.Context = context.TODO()
	return &wm
}

//GetAccountPendingTxCount
func (wm *WalletManager) GetAccountPendingTxCount(address string) (uint64, error) {

	//if wm.client == nil {
	//	return 0, fmt.Errorf("lisk API is not inited")
	//}
	//GetPendingAccountTransactionsByPubkey有bug
	//p := external.NewGetPendingAccountTransactionsByPubkeyParams().WithPubkey(address)
	//result, err := wm.Api.Node.External.GetPendingAccountTransactionsByPubkey(p)
	//if err != nil {
	//	return 0, err
	//}

	return uint64(0), nil
}

// BroadcastTransaction recalculates the transaction hash and sends the transaction to the node.
func (wm *WalletManager) BroadcastTransaction(txHex string) (string, error) {
	//txBytes, err := hex.DecodeString(txHex)
	//if err != nil {
	//	return "", fmt.Errorf("transaction decode failed, unexpected error: %v", err)
	//}
	//signedEncodedTx := aeternity.Encode(aeternity.PrefixTransaction, txBytes)
	// calculate the hash of the decoded txRLP
	//rlpTxHashRaw := owcrypt.Hash(txBytes, 32, owcrypt.HASH_ALG_BLAKE2B)
	//// base58/64 encode the hash with the th_ prefix
	//signedEncodedTxHash := aeternity.Encode(aeternity.PrefixTransactionHash, rlpTxHashRaw)

	// send it to the network
	//return postTransaction(wm.Api.Node, signedEncodedTx)
	return "", nil
}

func IntToBalance(balance int64) *big.Int {
	return big.NewInt(balance)
}
