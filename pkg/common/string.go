package common

import (
	"crypto/rand"
	"fmt"
	"log"
	"encoding/json"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func CreateUuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func String(ad interface{}) string {
	b, err := json.Marshal(ad)
	if err != nil {
		return fmt.Sprintf("%v", ad)
	}
	return string(b)
}

func HmacSha256(ad string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(ad))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5(str string) string {
	if len(str) == 0 {
		return ""
	}
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func StringToFloat64(str string) float64 {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return res
}

func Float64ToString(value float64) string {
	return fmt.Sprintf("%v", value)
}

func StringToInt(value string) int {
	res, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return res
}