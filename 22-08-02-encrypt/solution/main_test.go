package main

import (
	"encoding/base64"
	"os"
	"path"
	"testing"
)

func Test_SetRedis(t *testing.T) {
	setFileToRedis()
}

func Test_GetRedis(t *testing.T) {
	getFileFromRedisAndSetToTempDir()
}

func Test_base64Encoding(t *testing.T) {
	// read file to redis
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	split, file := path.Split(fileName)
	println(split, file)
	base64Str := base64.StdEncoding.EncodeToString(content)
	println(base64Str)
}
