package go_ompay

import (
	"fmt"
	"testing"
)

func TestFPXDeposit(t *testing.T) {
	//构造client
	cli := NewClient(nil,
		&OMPayInitParams{MerchantInfo{MERCHANT_ID,
			API_KEY,
			SECRET_KEY,
			WITHDRAW_AGENT_CODE,
			WITHDRAW_SECRET_KEY},
			DEPOSIT_URL,
			DEPOSIT_CALLBACK_URL,
			DEPOSIT_FE_CALLBACK_URL,
			DEPOSIT_FPX_URL,
			WITHDRAW_URL,
			WITHDRAW_CALLBACK_URL,
		})

	//获取拼凑的跳转地址
	cli.SetDebugModel(true)
	resp, err := cli.DepositFPXQR(GenDepositFPXRequestDemo())

	fmt.Printf("resp:%+v, err:%+v\n", resp, err)

}

//myr, idr,

// MYR 支持
// VND 支持
func GenDepositFPXRequestDemo() OMPayFPXDepositReq {
	return OMPayFPXDepositReq{
		SerialNo: "98905133224", //商户的单号
		Currency: "VND",
		Amount:   "1000000",
		//PayType:  "1004",
	}

}
