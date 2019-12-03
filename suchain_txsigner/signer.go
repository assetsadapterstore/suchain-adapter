package suchain_txsigner

var Default = &TransactionSigner{}

type TransactionSigner struct {
}

// SignSerialize 序列化签名
// required
func (singer *TransactionSigner) SignSerialize(sig []byte)[]byte {

	sb := sig[32:]
	rb := sig[:32]
	length := 6 + len(rb) + len(sb)
	b := make([]byte, length)
	b[0] = 0x30
	b[1] = byte(length - 2)
	b[2] = 0x02
	b[3] = byte(len(rb))
	offset := copy(b[4:], rb) + 4
	b[offset] = 0x02
	b[offset+1] = byte(len(sb))
	copy(b[offset+2:], sb)

	return b
}



