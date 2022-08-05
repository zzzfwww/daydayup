package main

import (
	"context"
	"encoding/base64"
	"log"
	"solution/crypto"

	"os"
	"path"
	"solution/redis"
)

const (
	fileName = "go.mod"
	key      = "123456781234567812345671"
)

var (
	ctx = context.Background()
	rc  = redis.NewRedisClint()
)

func setFileToRedis() {
	// read file to redis
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	split, file := path.Split(fileName)
	println(split, file)
	// rc.SetRedis(ctx, filepath, base64.StdEncoding.EncodeToString(content))
	base64Str := base64.StdEncoding.EncodeToString(content)
	encrypt := crypto.AesEncrypt(base64Str, key)
	rc.SetRedis(ctx, file, encrypt)
}
func getFileFromRedisAndSetToTempDir() {
	split, file := path.Split(fileName)
	println(split, file)
	writeDirFile := os.TempDir() + file
	println(writeDirFile)
	res := rc.GetRedis(ctx, fileName)
	decrypt := crypto.AesDecrypt(res, key)
	decodeString, err := base64.StdEncoding.DecodeString(decrypt)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = os.WriteFile(writeDirFile, decodeString, 0644)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	// rc.ExampleTest(ctx)
	split, file := path.Split(fileName)
	println(split, file)
	writeDirFile := os.TempDir() + file
	println(writeDirFile)
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.File(writeDirFile)
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
