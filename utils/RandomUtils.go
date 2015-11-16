package utils

import (
	"crypto/rand"
	"encoding/base32"
	"math/big"
	"strconv"
	"strings"
	"time"
)

type randomUtils struct {
}

var RandomUtils randomUtils = randomUtils{}

func (self *randomUtils) RandomString(length int) string {
	randomBytes := make([]byte, 1000)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}

func (self *randomUtils) RandomDigits(length int) string {
	var key string = ""
	for len(key) < length {
		key = key + strconv.Itoa(int(self.RandomDigit()))
	}
	return key
}

func (self *randomUtils) RandomDigit() uint32 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(9223372036854775800))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return uint32(n)
}

func (self *randomUtils) GetRandomFloat(bound int) float64 {
	ft := float64(self.RandomDigit()) / float64(2)
	for ft > float64(bound) {
		ft = ft / float64(2)
	}
	return ft
}

func (self *randomUtils) RandomFloat(bound int) string {
	return strconv.FormatFloat(self.GetRandomFloat(bound), 'f', 16, 64)
}

func (self *randomUtils) GetFileName() string {
	fileNamePrefix := "./log/apolloresponse"
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	timeStr = strings.Replace(timeStr, ":", "-", -1)
	return fileNamePrefix + timeStr
}
