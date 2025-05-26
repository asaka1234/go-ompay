package go_ompay

import (
	"github.com/asaka1234/go-ompay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID       string //貌似只是给deposit用的
	DepositApiKey    string
	DepositSecretKey string

	WithdrawAgentCode string
	WithdrawSecretKey string

	DepositUrl           string
	DepositCallbackUrl   string
	DepositFeCallbackUrl string

	WithdrawUrl         string
	WithdrawCallbackUrl string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantId string, depositApiKey string, depositSecretKey string, withdrawAgentCode, withdrawSecretKey, depositUrl, WithdrawUrl, depositCallbackUrl, WithdrawCallbackUrl, depositFeCallbackUrl string) *Client {
	return &Client{
		MerchantID:       merchantId,
		DepositApiKey:    depositApiKey,
		DepositSecretKey: depositSecretKey,

		WithdrawAgentCode: withdrawAgentCode,
		WithdrawSecretKey: withdrawSecretKey,

		WithdrawUrl:         WithdrawUrl,
		WithdrawCallbackUrl: WithdrawCallbackUrl,

		DepositUrl:           depositUrl,
		DepositCallbackUrl:   depositCallbackUrl,
		DepositFeCallbackUrl: depositFeCallbackUrl,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}

// 方便依据currency来切换 (MYR, IDR/ / SGD/ THB/ VND)
func (cli *Client) SetMerchantInfo(merchantId string, depositApiKey string, depositSecretKey string, withdrawAgentCode, withdrawSecretKey string) {
	cli.MerchantID = merchantId
	cli.DepositApiKey = depositApiKey
	cli.DepositSecretKey = depositSecretKey
	cli.WithdrawAgentCode = withdrawAgentCode
	cli.WithdrawSecretKey = withdrawSecretKey
}
