package util

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func GetRandom() int64 {
	return rand.Int63n(900000) + 100000
}

func GetCurrentTime() int64 {
	return time.Now().Unix()
}

func CalculateSignature(appKey string, random int64, timestamp int64) string {
	str := fmt.Sprintf("appkey=%s&random=%d&time=%d", appKey, random, timestamp)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func CalculateSignatureWithPhoneNumber(appKey string, random int64, timestamp int64, phoneNumber string) string {
	str := fmt.Sprintf("appkey=%s&random=%d&time=%d&mobile=%s", appKey, random, timestamp, phoneNumber)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func CalculateSignatureWithPhoneNumbers(appKey string, random int64, timestamp int64, phoneNumbers []string) string {

	phoneNumbersStr := ""
	if len(phoneNumbers) > 0 {
		phoneNumbersStr = phoneNumbers[0]
		for i := 1; i < len(phoneNumbers); i++ {
			phoneNumbersStr += ","
			phoneNumbersStr += phoneNumbers[i]
		}
	}

	str := fmt.Sprintf("appkey=%s&random=%d&time=%d&mobile=%s", appKey, random, timestamp, phoneNumbersStr)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func CalculateSignatureWithFid(appKey string, random int64, timestamp int64, fid string) string {
	str := fmt.Sprintf("appkey=%s&random=%d&time=%d&fid=%s", appKey, random, timestamp, fid)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func CalculateAuth(appKey string, random int64, timestamp int64, fileSha1Sum string) string {
	str := fmt.Sprintf("appkey=%s&random=%d&time=%d&content-sha1=%s", appKey, random, timestamp, fileSha1Sum)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func Sha1Sum(bytes []byte) string {
	h := sha1.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum([]byte(nil)))
}
