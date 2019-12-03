package suchain

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAddressDecoder_PublicKeyToAddress(t *testing.T) {
	pub, _ := hex.DecodeString("03aada4050092270cdfd931d1f507346e16883a51949be4d38542c4375d3a875f7")
	decoder := AddressDecoder{}
	addr, err := decoder.PublicKeyToAddress(pub, false)
	if err != nil {
		t.Errorf("PublicKeyToAddress error: %v", err)
		return
	}
	fmt.Println(addr)
}
