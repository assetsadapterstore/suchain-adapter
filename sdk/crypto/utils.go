// This file is part of SGU Go Crypto.
//
// (c) SGU Ecosystem <info@SGU.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package crypto

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/assetsadapterstore/suchain-adapter/sdk/crypto/base58"
)

func Byte2Hex(data byte) string {
	return fmt.Sprintf("%x", data)
}

func Hex2Byte(data []byte) string {
	return strings.ToLower(fmt.Sprintf("%X", data))
}

func HexEncode(data []byte) string {
	return hex.EncodeToString(data)
}

func HexDecode(data string) []byte {
	result, err := hex.DecodeString(data)

	if err != nil {
		log.Fatal(err.Error())
	}

	return result
}

func Base58Encode(data []byte) string {
	return base58.Encode(data)
}

func Base58Decode(data string) []byte {
	result, err := base58.Decode(data)

	if err != nil {
		log.Fatal(err.Error())
	}

	return result
}
