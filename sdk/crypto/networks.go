// This file is part of SGU Go Crypto.
//
// (c) SGU Ecosystem <info@SGU.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package crypto

import "time"

var (
	NETWORKS_MAINNET = &Network{
		Epoch:   time.Date(2017, 3, 21, 13, 00, 0, 0, time.UTC),
		Version: 23,
		Wif:     170,
	}
	NETWORKS_DEVNET = &Network{
		Epoch:   time.Date(2019, 10, 26, 06, 22, 16, 826, time.UTC),
		Version: 63,
		Wif:     198,
	}
	NETWORKS_TESTNET = &Network{
		Epoch:   time.Date(2017, 3, 21, 13, 00, 0, 0, time.UTC),
		Version: 23,
		Wif:     186,
	}
)
