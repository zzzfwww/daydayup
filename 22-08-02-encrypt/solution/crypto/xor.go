package crypto

import (
	"encoding/base64"
)

func XorEncrypt(orig string, key byte) string {
	// 转成字节数组
	origData := []byte(orig)
	// 创建数组
	crypto := make([]byte, len(origData))
	for ind := range origData {
		crypto[ind] = origData[ind] ^ key
	}
	return base64.StdEncoding.EncodeToString(crypto)

}

func XorDecrypt(cryted string, key byte) string {
	cryptoByte, _ := base64.StdEncoding.DecodeString(cryted)
	orig := make([]byte, len(cryptoByte))
	for ind := range cryptoByte {
		orig[ind] = cryptoByte[ind] ^ key
	}

	return string(orig)
}
