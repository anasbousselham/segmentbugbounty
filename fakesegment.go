package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"

	"github.com/icrowley/fake"
)

const (
	keyList string = "abcdefghijklmnopqrstuvwxyzABCDEFHFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

func main() {
	argsWithProg := os.Args[1]
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
	}

	f.Close()
	url := "https://api.segment.io/v1/p"
	content, err := ioutil.ReadFile("keygen")
	if err != nil {
		log.Fatal(err)
	}
	name := fake.FirstName()
	//product := fake.Product()
	ip := fake.IPv4()
	var jsonStr = []byte(`{"timestamp":"2021-01-07T12:51:59.412Z",
	"context":{"page":{"path":"/","referrer":"","search":"","title":"` + string(name) + `",
	"url":"https://bousselham.com/"},
	"userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36","locale":"en-US",
	"library":{"name":"analytics.js","version":"4.0.4"}},"integrations":{},
	"properties":{"path":"/","referrer":"",
	"search":"",
	"title":"` + string(name) + `",
	"url":"https://bousselham.com/",
	"ip_address":"` + string(ip) + `"
	},
	"messageId":"ajs-` + string(content) + `",
	"anonymousId":"3d8bc2e8-ef39-4ffb-adbb-296adbbc799c",
	"type":"page",
	"writeKey":"` + argsWithProg + `",
	"userId":null,"sentAt":"2021-01-07T12:51:59.414Z",
	"_metadata":{"bundled":["Segment.io"],"unbundled":[]}}`)
	//debug --> fmt.Printf("%s", jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Origin", "https://bousselham.com/")
	req.Header.Set("Referer", "https://bousselham.com/")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//	body, _ := ioutil.ReadAll(resp.Body)
	//	fmt.Println("response Body:", string(body))
}
