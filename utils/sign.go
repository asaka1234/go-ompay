package utils

import (
	"crypto/md5"
	"fmt"
)

// MD5(serialNo + {API KEY} + {Secret KEY} + amount)
func SignWithAmount(serialNo, amount, apiKey, apiSecret string) string {
	// Generate MD5 hash from the concatenated string
	hashStr := fmt.Sprintf("%s%s%s%s", serialNo, apiKey, apiSecret, amount)
	hash := md5.Sum([]byte(hashStr))
	md5Token := fmt.Sprintf("%x", hash) // Convert to hex string
	return md5Token
}

// MD5(serialNo + {API KEY} + {Secret KEY})
func SignWithoutAmount(serialNo, apiKey, apiSecret string) string {
	// Generate MD5 hash from the concatenated string
	hashStr := fmt.Sprintf("%s%s%s", serialNo, apiKey, apiSecret)
	hash := md5.Sum([]byte(hashStr))
	md5Token := fmt.Sprintf("%x", hash) // Convert to hex string
	return md5Token
}

// 验证签名
func VerifySignWithoutAmount(serialNo, apiKey, apiSecret string, signKey string) bool {

	// Generate current signature
	currentKey := SignWithoutAmount(serialNo, apiKey, apiSecret)

	// Compare the signatures
	return signKey == currentKey
}
