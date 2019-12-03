// This file is part of SGU Go Client.
//
// (c) SGU Ecosystem <info@SGU.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type Transaction struct {
	Address       string            `json:"address,omitempty"`
	Id            string            `json:"id,omitempty"`
	BlockId       string            `json:"blockId,omitempty"`
	BlockHeight   int64             `json:"-"`
	Type          uint64              `json:"type,omitempty"`
	Amount        string            `json:"amount,omitempty"`
	Fee           string            `json:"fee,omitempty"`
	Sender        string            `json:"sender,omitempty"`
	Recipient     string            `json:"recipient,omitempty"`
	Signature     string            `json:"signature,omitempty"`
	VendorField   string            `json:"vendorField,omitempty"`
	Asset         *TransactionAsset `json:"asset,omitempty"`
	Confirmations uint16            `json:"confirmations,omitempty"`
	Timestamp     Timestamp         `json:"timestamp,omitempty"`
}

type Transactions struct {
	Meta Meta          `json:"meta,omitempty"`
	Data []Transaction `json:"data,omitempty"`
}

type GetTransaction struct {
	Data Transaction `json:"data,omitempty"`
}

type GetCreateTransaction struct {
	Data CreateTransaction `json:"data,omitempty"`
}

type TransactionTypes struct {
	Data map[string]byte `json:"data,omitempty"`
}

type Timestamp struct {
	Epoch int32  `json:"epoch,omitempty"`
	Unix  int32  `json:"unix,omitempty"`
	Human string `json:"human,omitempty"`
}

type CreateTransaction struct {
	Accept  []string `json:"accept,omitempty"`
	Excess  []string `json:"excess,omitempty"`
	Invalid []string `json:"invalid,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////
// TRANSACTION ASSETS //////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

type TransactionAsset struct {
	Votes          []string                          `json:"votes,omitempty"`
	Signature      *SecondSignatureRegistrationAsset `json:"signature,omitempty"`
	Delegate       *DelegateAsset                    `json:"publicKey,omitempty"`
	MultiSignature *MultiSignatureRegistrationAsset  `json:"multisignature,omitempty"`
	Ipfs           *IpfsAsset                        `json:"ipfs,omitempty"`
	Payments       []*MultiPaymentAsset              `json:"payments,omitempty"`
}

type SecondSignatureRegistrationAsset struct {
	PublicKey string `json:"publicKey,omitempty"`
}

type DelegateAsset struct {
	Username string `json:"username,omitempty"`
}

type MultiSignatureRegistrationAsset struct {
	Min       byte     `json:"min,omitempty"`
	Keysgroup []string `json:"keysgroup,omitempty"`
	Lifetime  byte     `json:"lifetime,omitempty"`
}

type IpfsAsset struct {
	Dag string `json:"dag,omitempty"`
}

type MultiPaymentAsset struct {
	Amount      uint64 `json:"amount,omitempty"`
	RecipientId string `json:"recipientId,omitempty"`
}
