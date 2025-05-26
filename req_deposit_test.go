package go_ompay

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {
	//构造client
	cli := NewClient(nil,
		MERCHANT_ID,
		API_KEY,
		SECRET_KEY,
		WITHDRAW_AGENT_CODE,
		WITHDRAW_SECRET_KEY,
		DEPOSIT_URL,
		WITHDRAW_URL,
		DEPOSIT_CALLBACK_URL,
		WITHDRAW_CALLBACK_URL,
		DEPOSIT_FE_CALLBACK_URL)

	//获取拼凑的跳转地址
	urlString := cli.Deposit(GenDepositRequestDemo())

	fmt.Printf("resp:%+v\n", urlString)

}

func GenDepositRequestDemo() OMPayDepositReq {
	return OMPayDepositReq{
		SerialNo:          "11234",     //商户的单号  //1
		Currency:          "VND",       //1
		Amount:            "100000.00", //Return URL after the payment is done.
		ClientAccountName: "你好",        //Client's Registered Full name in account (KYC)
	}

}
