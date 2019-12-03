package suchain_addrec

import (
	"encoding/hex"
	"fmt"
	"testing"
)


func TestAddressDecoder_PublicKeyToAddress(t * testing.T){
	pub, _ := hex.DecodeString("03aada4050092270cdfd931d1f507346e16883a51949be4d38542c4375d3a875f7")
	address := GetAddressFromPublicKey(pub)
	fmt.Println(address)
}