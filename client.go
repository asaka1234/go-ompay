package go_ompay

import (
	"github.com/asaka1234/go-ompay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID string
	ApiKey     string
	ApiSecret  string

	DepositUrl           string
	DepositCallbackUrl   string
	DepositFeCallbackUrl string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantId string, apiKey string, apiSecret string, depositUrl, depositCallbackUrl, depositFeCallbackUrl string) *Client {
	return &Client{
		MerchantID: merchantId,
		ApiKey:     apiKey,
		ApiSecret:  apiSecret,

		DepositUrl:           depositUrl,
		DepositCallbackUrl:   depositCallbackUrl,
		DepositFeCallbackUrl: depositFeCallbackUrl,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
