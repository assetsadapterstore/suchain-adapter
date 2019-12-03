package openwtester

import (
	"github.com/assetsadapterstore/suchain-adapter/suchain"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	// openw.RegAssets(eosio.Symbol, eosio.NewWalletManager(nil))


	openw.RegAssets(suchain.Symbol, suchain.NewWalletManager())
}
