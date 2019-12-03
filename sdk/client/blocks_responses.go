// This file is part of SGU Go Client.
//
// (c) SGU Ecosystem <info@SGU.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import "github.com/blocktree/openwallet/openwallet"

type BlockForged struct {
	Reward int64 `json:"reward,omitempty"`
	Fee    int64 `json:"fee,omitempty"`
	Total  int64 `json:"total,omitempty"`
}

type BlockPayload struct {
	Hash   string `json:"hash,omitempty"`
	Length byte   `json:"length,omitempty"`
}

type BlockGenerator struct {
	Username  string `json:"username,omitempty"`
	Address   string `json:"address,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
}

type Block struct {
	Id           string         `json:"id,omitempty"`
	Version      byte           `json:"version,omitempty"`
	Height       int64          `json:"height,omitempty"  storm:"id"`
	Previous     string         `json:"previous,omitempty"`
	//Forged       BlockForged    `json:"forged,omitempty"`
	//Payload      BlockPayload   `json:"payload,omitempty"`
	//Generator    BlockGenerator `json:"generator,omitempty"`
	Signature    string         `json:"signature,omitempty"`
	Transactions byte           `json:"transactions,omitempty"`
	Timestamp    Timestamp      `json:"timestamp,omitempty"`
}


//BlockHeader 区块链头
func (b *Block) BlockHeader(symbol string) *openwallet.BlockHeader {

	obj := openwallet.BlockHeader{}
	//解析json
	obj.Hash = b.Id
	obj.Previousblockhash = b.Previous
	obj.Height = uint64(b.Height)
	obj.Version = uint64(b.Version)
	obj.Time = uint64(b.Timestamp.Unix)
	obj.Symbol = symbol

	return &obj
}


type Blocks struct {
	Meta Meta    `json:"meta,omitempty"`
	Data []Block `json:"data,omitempty"`
}

type GetBlock struct {
	Data Block `json:"data,omitempty"`
}

type GetBlockTransactions struct {
	Meta Meta          `json:"meta,omitempty"`
	Data []Transaction `json:"data,omitempty"`
}
