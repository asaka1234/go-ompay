package go_ompay

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func TestDeposit(t *testing.T) {
	//构造client
	vLog := VLog{}
	cli := NewClient(vLog,
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
	//urlString := cli.Deposit(GenDepositRequestDemo())

	//fmt.Printf("resp:%+v\n", urlString)

	// fpx发请求给psp三方
	respfpx, err := cli.DepositFPXQR(OMPayFPXDepositReq{
		Currency: "MYR",
		Amount:   cast.ToString(3000), //这里改为上游入金的流水号
		SerialNo: "20253484848483",    //outNo
	})
	if err != nil {
		fmt.Printf("respfpx err:%+v\n", err)
	}
	fmt.Printf("respfpx:%+v\n", respfpx)

}

func GenDepositRequestDemo() OMPayDepositReq {
	return OMPayDepositReq{
		SerialNo:          "789012",  //商户的单号  //1
		Currency:          "THB",     //1
		Amount:            "1000000", //Return URL after the payment is done.
		ClientAccountName: "hello",   //Client's Registered Full name in account (KYC)
	}

}
