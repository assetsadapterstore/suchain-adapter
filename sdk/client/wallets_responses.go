// This file is part of SGU Go Client.
//
// (c) SGU Ecosystem <info@SGU.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type Wallet struct {
	Address    string `json:"address,omitempty"`
	PublicKey  string `json:"publicKey,omitempty"`
	Balance    string  `json:"balance,omitempty"`
	IsDelegate bool   `json:"isDelegate,omitempty"`
}

type Wallets struct {
	Meta Meta     `json:"meta,omitempty"`
	Data []Wallet `json:"data,omitempty"`
}

type GetWallet struct {
	Data Wallet `json:"data,omitempty"`
}
