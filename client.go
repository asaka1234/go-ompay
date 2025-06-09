package go_ompay

import (
	"github.com/asaka1234/go-ompay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params OMPayInitParams

	ryClient  *resty.Client
	debugMode bool
	logger    utils.Logger
}

func NewClient(logger utils.Logger, params OMPayInitParams) *Client {
	return &Client{
		Params: params,

		ryClient:  resty.New(), //client实例
		debugMode: false,
		logger:    logger,
	}
}

// 方便依据currency来切换 (MYR, IDR/ / SGD/ THB/ VND)
func (cli *Client) SetMerchantInfo(merchant MerchantInfo) {
	cli.Params.MerchantInfo = merchant
}

func (cli *Client) SetDebugModel(debugMode bool) {
	cli.debugMode = debugMode
}
