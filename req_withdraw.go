package go_ompay

import (
	"github.com/asaka1234/go-ompay/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"net/url"
)

// https://api.doitwallet.asia/Documents/PayoutAPI.pdf
func (cli *Client) Withdraw(req OMPayWithdrawalReq) string {

	rawURL := cli.DepositUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["AgentCode"] = cli.WithdrawAgentCode     //系统分配
	params["CallbackURL"] = cli.WithdrawCallbackUrl //回调地址

	//签名
	signStr := utils.SignWithdrawWithUserRef(req.UserRef, cli.WithdrawAgentCode, cli.WithdrawSecretKey)
	params["Token"] = signStr

	//http://<domain>/Merchant/Pay?merchantCode={Merchant Id}&serialNo={Your
	//Transaction id} &currency={Currency}&amount={Amount}&returnUrl={Return URL}
	//&notifyUrl&={Callback URL} &token={MD5 token}

	//构造url
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	// 2. 设置查询参数
	q := u.Query()
	for key, value := range params {
		q.Add(key, cast.ToString(value))
	}
	u.RawQuery = q.Encode()

	return u.String()

}
