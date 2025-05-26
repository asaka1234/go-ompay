package go_ompay

import (
	"errors"
	"github.com/asaka1234/go-ompay/utils"
	"github.com/mitchellh/mapstructure"
)

// POST请求传过来的. 但是所有参数都在query里
// https://{your_callback_url}/?Token=***&TransactionId=ABC123456&StatusDesc=Complete
func (cli *Client) WithdrawCallback(req OmPayWithdrawalCallbackReq, processor func(OmPayWithdrawalCallbackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySignWithdrawWithTransId(req.TransactionId, cli.WithdrawAgentCode, cli.WithdrawSecretKey, req.Token)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}

	//开始处理
	return processor(req)
}
