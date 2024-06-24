package adr

import "encoding/base64"

const (
	base64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"
)

var coder = base64.NewEncoding(base64Table)

//编码
func Base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

//解码
func Base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}