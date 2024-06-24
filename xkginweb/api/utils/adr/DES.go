package adr

import (
	"fmt"
	"github.com/agclqq/goencryption"
)

var key = []byte("kuangxx.")
var iv = []byte("kuangxx.")

func DesCBCEncrypt(src string) string {
	cryptText, err := goencryption.DesCBCPkcs7Encrypt([]byte(src), key, iv)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	out := goencryption.Base64Encode(cryptText)
	return out
}

func DesCBCDecrypt(out string) string {
	cText, err := goencryption.Base64Decode(out)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	decrypt, _ := goencryption.DescCBCPkcs7Decrypt(cText, key, iv)
	return string(decrypt)
}
