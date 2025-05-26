package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// MD5(serialNo + {API KEY} + {Secret KEY} + amount)
func SignDepositWithAmount(serialNo, amount, apiKey, apiSecret string) string {
	// Generate MD5 hash from the concatenated string
	hashStr := fmt.Sprintf("%s%s%s%s", serialNo, apiKey, apiSecret, amount)
	hash := md5.Sum([]byte(hashStr))
	md5Token := fmt.Sprintf("%x", hash) // Convert to hex string
	return md5Token
}

// MD5(serialNo + {API KEY} + {Secret KEY})
func SignDepositWithoutAmount(serialNo, apiKey, apiSecret string) string {
	// Generate MD5 hash from the concatenated string
	hashStr := fmt.Sprintf("%s%s%s", serialNo, apiKey, apiSecret)
	hash := md5.Sum([]byte(hashStr))
	md5Token := fmt.Sprintf("%x", hash) // Convert to hex string
	return md5Token
}

// 验证签名
// MD5(serialNo + {API KEY} + {Secret KEY})
func VerifySignDepositWithoutAmount(serialNo, apiKey, apiSecret string, signKey string) bool {

	// Generate current signature
	currentKey := SignDepositWithoutAmount(serialNo, apiKey, apiSecret)

	// Compare the signatures
	return signKey == currentKey
}

//---------------------

// MD5({Agent Code} + UserRef.ToUpper() + {Secret KEY})
func SignWithdrawWithUserRef(userRef, agentCode, secretKey string) string {
	// Generate MD5 hash from the concatenated string
	hashStr := fmt.Sprintf("%s%s%s", agentCode, strings.ToUpper(userRef), secretKey)
	hash := md5.Sum([]byte(hashStr))
	md5Token := fmt.Sprintf("%x", hash) // Convert to hex string
	return md5Token
}

// MD5(AgentCode.ToUpper() + TransactionId.ToUpper() + API_SECRET_KEY)
func SignWithdrawWithTransId(transactionId, agentCode, secretKey string) string {
	// Generate MD5 hash from the concatenated string
	hashStr := fmt.Sprintf("%s%s%s", strings.ToUpper(agentCode), strings.ToUpper(transactionId), secretKey)
	hash := md5.Sum([]byte(hashStr))
	md5Token := fmt.Sprintf("%x", hash) // Convert to hex string
	return md5Token
}

func VerifySignWithdrawWithTransId(transactionId, agentCode, secretKey string, signKey string) bool {

	// Generate current signature
	currentKey := SignWithdrawWithTransId(transactionId, agentCode, secretKey)

	// Compare the signatures
	return signKey == currentKey
}
