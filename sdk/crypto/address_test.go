// This file is part of SGU Go Crypto.
//
// (c) SGU Ecosystem <info@SGU.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddressFromPassphrase(t *testing.T) {
	fixture := GetIdentityFixture()

	address, _ := AddressFromPassphrase(fixture.Passphrase)

	assert := assert.New(t)
	assert.Equal(fixture.Data.Address, address)
}

func TestValidateAddress(t *testing.T) {
	//fixture := GetIdentityFixture()

	assert := assert.New(t)
	resutl ,_:= ValidateAddress("STB9Cs3cCwoQ11TidumhAdcmvhLTZfHFUg")
	assert.True(resutl)
}
