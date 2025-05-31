package go_ompay

import (
	"fmt"
	"testing"
)

func TestWithdraw(t *testing.T) {
	//构造client
	cli := NewClient(nil,
		OMPayInitParams{MerchantInfo{MERCHANT_ID,
			API_KEY,
			SECRET_KEY,
			WITHDRAW_AGENT_CODE,
			WITHDRAW_SECRET_KEY},
			DEPOSIT_URL,
			DEPOSIT_CALLBACK_URL,
			DEPOSIT_FE_CALLBACK_URL,
			WITHDRAW_URL,
			WITHDRAW_CALLBACK_URL})

	//获取拼凑的跳转地址
	resp, _ := cli.Withdraw(GenWithdrawRequestDemo())

	fmt.Printf("resp:%+v\n", resp)

}

func GenWithdrawRequestDemo() OMPayWithdrawalReq {
	return OMPayWithdrawalReq{
		UserRef:           "11234", //商户的单号  //1
		TransactionId:     "3444",
		FullName:          "你好",
		AccountNo:         "111",
		BankCode:          "MBB",
		BankRegister:      "哈哈",
		BankRegisterState: "11",
		BankRegisterCity:  "上海",
		Currency:          "VND", //1
		Amount:            100000.00,
	}

}
