package go_ompay

import (
	"crypto/tls"
	"github.com/asaka1234/go-ompay/utils"
	"github.com/mitchellh/mapstructure"
)

// 用fpx和qr code支付
func (cli *Client) DepositFPXQR(req OMPayFPXDepositReq) (*OMPayFPXDepositResponse, error) {

	rawURL := cli.Params.FPXDepositUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["merchantCode"] = cli.Params.MerchantId //1
	params["returnUrl"] = cli.Params.DepositFeBackUrl
	params["notifyUrl"] = cli.Params.DepositBackUrl //回调地址

	//签名
	signStr := utils.SignDepositWithAmount(req.SerialNo, req.Amount, cli.Params.DepositApiKey, cli.Params.DepositSecretKey)
	params["token"] = signStr

	//返回值会放到这里
	var result OMPayFPXDepositResponse

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getFPXHeaders()).
		SetBody(params).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
