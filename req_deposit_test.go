package go_ompay

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {
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
	urlString := cli.Deposit(GenDepositRequestDemo())

	fmt.Printf("resp:%+v\n", urlString)

}

func GenDepositRequestDemo() OMPayDepositReq {
	return OMPayDepositReq{
		SerialNo:          "789012",  //商户的单号  //1
		Currency:          "THB",     //1
		Amount:            "1000000", //Return URL after the payment is done.
		ClientAccountName: "hello",   //Client's Registered Full name in account (KYC)
	}

}
