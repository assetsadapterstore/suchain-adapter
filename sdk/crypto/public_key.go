// This file is part of SGU Go Crypto.
//
// (c) SGU Ecosystem <info@SGU.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package crypto

import (
	"github.com/assetsadapterstore/suchain-adapter/sdk/crypto/base58"
	"github.com/btcsuite/btcd/btcec"
	"golang.org/x/crypto/ripemd160"
	"time"
)

func PublicKeyFromPassphrase(passphrase string) (*PublicKey, error) {
	privateKey, err := PrivateKeyFromPassphrase(passphrase)

	if err != nil {
		return nil, err
	}

	return privateKey.PublicKey, nil
}

func PublicKeyFromHex(publicKeyHex string) (*PublicKey, error) {
	publicKey, err := PublicKeyFromBytes(HexDecode(publicKeyHex))

	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func PublicKeyFromBytes(bytes []byte) (*PublicKey, error) {


	SetNetwork(&Network{
		Epoch:   time.Date(2019, 10, 26, 06, 22, 16, 826, time.UTC),
		Version: 63,
		Wif:     198,
	})
	publicKey, err := btcec.ParsePubKey(bytes, btcec.S256())

	if err != nil {
		return nil, err
	}

	isCompressed := false

	if len(bytes) == btcec.PubKeyBytesLenCompressed {
		isCompressed = true
	}


	return &PublicKey{
		PublicKey:    publicKey,
		isCompressed: isCompressed,
		Network:      GetNetwork(),
	}, nil
}

////////////////////////////////////////////////////////////////////////////////
// ADDRESS COMPUTATION /////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

func (publicKey *PublicKey) ToHex() string {
	return HexEncode(publicKey.Serialize())
}

func (publicKey *PublicKey) ToAddress() string {
	ripeHashedBytes := publicKey.AddressBytes()
	ripeHashedBytes = append(ripeHashedBytes, 0x0)
	copy(ripeHashedBytes[1:], ripeHashedBytes[:len(ripeHashedBytes)-1])
	ripeHashedBytes[0] = publicKey.Network.Version

	return base58.Encode(ripeHashedBytes)
}

func (publicKey *PublicKey) Serialize() []byte {
	if publicKey.isCompressed {
		return publicKey.SerializeCompressed()
	}

	return publicKey.SerializeUncompressed()
}

func (publicKey *PublicKey) AddressBytes() []byte {
	hash := ripemd160.New()
	_, _ = hash.Write(publicKey.Serialize())

	return hash.Sum(nil)
}
