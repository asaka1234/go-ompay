package go_ompay

import (
	"github.com/asaka1234/go-ompay/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"net/url"
)

// 构造一个支付地址, 随后让前端Get打开这个地址即可跳转到psp三方收银台(是一个qrcode扫描支付的页面)
func (cli *Client) Deposit(req OMPayDepositReq) string {

	rawURL := cli.Params.DepositUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["merchantCode"] = cli.Params.MerchantId //1
	params["returnUrl"] = cli.Params.DepositFeBackUrl
	params["notifyUrl"] = cli.Params.DepositBackUrl //回调地址

	//签名
	signStr := utils.SignDepositWithAmount(req.SerialNo, req.Amount, cli.Params.DepositApiKey, cli.Params.DepositSecretKey)
	params["token"] = signStr

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
