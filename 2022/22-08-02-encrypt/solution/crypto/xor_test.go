package crypto

import (
	"fmt"
	"testing"
)

func TestXor(t *testing.T) {
	orig := "hello world"
	key := byte('x')

	fmt.Println("原文：", orig)
	encryptData := XorEncrypt(orig, key)
	fmt.Println("密文：", encryptData)

	decryptCode := XorDecrypt(encryptData, key)
	fmt.Println("解密结果：", decryptCode)
}
