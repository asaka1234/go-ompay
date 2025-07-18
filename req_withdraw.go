package go_ompay

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-ompay/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
)

// https://api.doitwallet.asia/Documents/PayoutAPI.pdf
func (cli *Client) Withdraw(req OMPayWithdrawalReq) (*OMPayWithdrawalResp, error) {

	rawURL := cli.Params.WithdrawUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["AgentCode"] = cli.Params.WithdrawAgentCode //系统分配
	params["CallbackURL"] = cli.Params.WithdrawBackUrl //回调地址

	//签名
	signStr := utils.SignWithdrawWithUserRef(req.UserRef, cli.Params.WithdrawAgentCode, cli.Params.WithdrawSecretKey)
	params["Token"] = signStr

	//----------------------
	var result OMPayWithdrawalResp

	resp1, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetFormData(utils.ConvertToStringMap(params)).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp1))
	cli.logger.Infof("PSPResty#ompay#withdraw->%+v", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp1.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp1.StatusCode())
	}

	if resp1.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%+v", resp1.Error())
	}

	if result.HasError {
		result.ErrorMessage = result.Info
		result.ErrorCode = 108 //TODO 这里要做一下映射
		result.Success = false
	}

	return &result, nil

}
