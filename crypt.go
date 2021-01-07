package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const (
	keyList string = "abcdefghijklmnopqrstuvwxyzABCDEFHFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

func main() {
	size := "32"
	strLen, _ := strconv.Atoi(size)
	filename := "keygen"
	os.Create(filename)
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0777)
	for key := 1; key <= strLen; key++ {
		res, _ := rand.Int(rand.Reader, big.NewInt(64))
		keyGen := keyList[res.Int64()]
		stringGen := fmt.Sprintf("%c", keyGen)
		f.Write([]byte(stringGen))
		fmt.Printf("%v", stringGen)
	}
	f.Close()
}
