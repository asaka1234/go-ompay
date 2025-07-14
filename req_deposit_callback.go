package go_ompay

import (
	"errors"
	"github.com/asaka1234/go-ompay/utils"
	"github.com/mitchellh/mapstructure"
)

// banktransfer和fpx都是同一个回调
// https://{your notify or callback url}?Info=Approved&MerchantCode=60-00000100-0000123&SerialNo=AA0123456&CurrencyCode=MYR&Amount=500.00&Status=1&Token=***
func (cli *Client) DepositCallback(req OmPayDepositCallbackReq, processor func(OmPayDepositCallbackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySignDepositWithoutAmount(req.SerialNo, cli.Params.DepositApiKey, cli.Params.DepositSecretKey, req.Token)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}
	if req.MerchantCode != cli.Params.MerchantId {
		return errors.New("merchanID is wrong!")
	}

	//开始处理
	return processor(req)
}
