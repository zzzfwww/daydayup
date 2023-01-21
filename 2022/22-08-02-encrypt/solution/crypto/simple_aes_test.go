package crypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAesDecrypt(t *testing.T) {
	orig := "hello world"
	key := "123456781234567812345678"
	fmt.Println("原文：", orig)

	encryptCode := AesEncrypt(orig, key)
	fmt.Println("密文：", encryptCode)

	decryptCode := AesDecrypt(encryptCode, key)
	fmt.Println("解密结果：", decryptCode)
}

func TestBase64(t *testing.T) {
	orig := "hello world"
	fmt.Println("原文：", orig)

	encryptCode := base64.StdEncoding.EncodeToString([]byte(orig))
	fmt.Println("密文：", encryptCode)

	decryptCode, err := base64.StdEncoding.DecodeString(encryptCode)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println("解密结果：", string(decryptCode))
}
