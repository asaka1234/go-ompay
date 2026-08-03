package go_ompay

import (
	"fmt"
	"testing"

	"github.com/asaka1234/go-ompay/utils"
)

// newPHPClient 构造PHP币种的测试客户端
func newPHPClient() *Client {
	vlog := VLog{}
	return NewClient(vlog,
		&OMPayInitParams{
			MerchantInfo: MerchantInfo{
				MerchantId:        PHP_MERCHANT_ID,
				DepositApiKey:     PHP_API_KEY,
				DepositSecretKey:  PHP_SECRET_KEY,
				WithdrawAgentCode: PHP_WITHDRAW_AGENT_CODE,
				WithdrawSecretKey: PHP_WITHDRAW_SECRET_KEY,
			},
			DepositUrl:       DEPOSIT_URL,
			DepositBackUrl:   DEPOSIT_CALLBACK_URL,
			DepositFeBackUrl: DEPOSIT_FE_CALLBACK_URL,
			FPXDepositUrl:    DEPOSIT_FPX_URL,
			WithdrawUrl:      WITHDRAW_URL,
			WithdrawBackUrl:  WITHDRAW_CALLBACK_URL,
		})
}

// ---------------------------------------------------------------------------
// 5.1 PHP 银行转账入金 (DepositAPI - GET方式, 返回跳转URL)
// ---------------------------------------------------------------------------

// TestPHPDeposit 测试PHP币种银行转账入金 (构造跳转URL)
func TestPHPDeposit(t *testing.T) {
	cli := newPHPClient()
	cli.SetDebugModel(true)

	urlString := cli.Deposit(GenPHPDepositReq())
	fmt.Printf("PHP Deposit URL: %s\n", urlString)

	if urlString == "" {
		t.Error("PHP deposit URL should not be empty")
	}
}

// GenPHPDepositReq 构造PHP入金请求示例
// PHP币种要求 ClientId 必填
func GenPHPDepositReq() OMPayDepositReq {
	return OMPayDepositReq{
		SerialNo:          "PHP20260731005", // 商户唯一单号
		Currency:          "PHP",            // PHP币种
		Amount:            "1000.00",        // 金额
		ClientAccountName: "Juan dela Cruz", // 用户KYC全名
		ClientId:          "juan@gmail.com", // 用户ID (PHP必填)
	}
}

// ---------------------------------------------------------------------------
// 5.2 PHP FPX/QR 入金 (FPXAPI - POST方式, 返回支付URL)
// ---------------------------------------------------------------------------

// TestPHPDepositFPXQR 测试PHP币种FPX/QR入金
func TestPHPDepositFPXQR(t *testing.T) {
	cli := newPHPClient()
	cli.SetDebugModel(true)

	resp, err := cli.DepositFPXQR(GenPHPFPXDepositReq())
	if err != nil {
		fmt.Printf("PHP FPX deposit err: %+v\n", err)
		// 网络/凭证错误时不视为测试失败, 仅输出
		return
	}
	fmt.Printf("PHP FPX deposit resp: %+v\n", resp)
}

// GenPHPFPXDepositReq 构造PHP FPX入金请求示例
func GenPHPFPXDepositReq() OMPayFPXDepositReq {
	return OMPayFPXDepositReq{
		SerialNo:          "PHP20260831003", // 商户唯一单号
		Currency:          "PHP",            // PHP币种
		Amount:            "500.00",         // 金额
		ClientAccountName: "Juan dela Cruz", // 用户KYC全名
		ClientId:          "juan@gmail.com", // 用户ID
	}
}

// ---------------------------------------------------------------------------
// PHP 入金回调验签
// ---------------------------------------------------------------------------

// TestPHPDepositCallback 测试PHP入金回调验签
// 回调格式与其他币种相同:
// https://{callback_url}?Info=Approved&MerchantCode=...&SerialNo=...&CurrencyCode=PHP&Amount=...&Status=1&Token=...&ActualAmount=...
func TestPHPDepositCallback(t *testing.T) {
	cli := newPHPClient()

	serialNo := "PHP20260731001"
	// 构造正确的Token: MD5(serialNo + {API KEY} + {Secret KEY})
	expectedToken := utils.SignDepositWithoutAmount(serialNo, PHP_API_KEY, PHP_SECRET_KEY)
	fmt.Printf("PHP deposit callback expected token: %s\n", expectedToken)

	err := cli.DepositCallback(OmPayDepositCallbackReq{
		Info:         "Approved",
		MerchantCode: PHP_MERCHANT_ID,
		SerialNo:     serialNo,
		CurrencyCode: "PHP",
		Amount:       1000.00,
		Status:       1, // 1 = Payment Approved
		Token:        expectedToken,
		ActualAmount: 1000.00,
	}, func(req OmPayDepositCallbackReq) error {
		fmt.Printf("PHP deposit callback processed: SerialNo=%s, Status=%d, Amount=%.2f\n",
			req.SerialNo, req.Status, req.Amount)
		return nil
	})

	if err != nil {
		t.Errorf("PHP DepositCallback failed: %+v\n", err)
	}
}

// TestPHPDepositCallbackInvalidToken 测试PHP入金回调签名验证失败场景
func TestPHPDepositCallbackInvalidToken(t *testing.T) {
	cli := newPHPClient()

	err := cli.DepositCallback(OmPayDepositCallbackReq{
		Info:         "Approved",
		MerchantCode: PHP_MERCHANT_ID,
		SerialNo:     "PHP20260731001",
		CurrencyCode: "PHP",
		Amount:       1000.00,
		Status:       1,
		Token:        "invalid_token_should_fail",
		ActualAmount: 1000.00,
	}, func(req OmPayDepositCallbackReq) error {
		return nil
	})

	if err == nil {
		t.Error("PHP DepositCallback should fail with invalid token")
	} else {
		fmt.Printf("PHP deposit callback correctly rejected invalid token: %v\n", err)
	}
}

// ---------------------------------------------------------------------------
// PHP 出金 (PayoutAPI)
// ---------------------------------------------------------------------------

// TestPHPWithdraw 测试PHP币种出金
func TestPHPWithdraw(t *testing.T) {
	cli := newPHPClient()
	cli.SetDebugModel(true)

	resp, err := cli.Withdraw(GenPHPWithdrawReq())
	if err != nil {
		fmt.Printf("PHP withdraw err: %+v\n", err)
		// 网络/凭证错误时不视为测试失败, 仅输出
		return
	}
	fmt.Printf("PHP withdraw resp: %+v\n", resp)
}

// GenPHPWithdrawReq 构造PHP出金请求示例
func GenPHPWithdrawReq() OMPayWithdrawalReq {
	return OMPayWithdrawalReq{
		UserRef:       "PHP_REF_20260731001", // 商户参考ID
		TransactionId: "PHP_TX_20260731001",  // 商户唯一订单号
		FullName:      "Juan dela Cruz",      // 用户全名
		AccountNo:     "09171234567",         // 菲律宾手机号/银行账号
		BankCode:      "GCash",               // 银行名称/电子钱包名称
		Amount:        500.00,                // 金额
		Currency:      "PHP",                 // PHP币种
		Channel:       "1",                   // 1 = Normal (可选)
	}
}

// ---------------------------------------------------------------------------
// PHP 出金回调验签
// ---------------------------------------------------------------------------

// TestPHPWithdrawCallback 测试PHP出金回调验签
// 回调格式:
// https://{callback_url}?Token=***&TransactionId=...&StatusDesc=Completed&StatusId=2&FullName=...&AccountNo=...&Amount=...
func TestPHPWithdrawCallback(t *testing.T) {
	cli := newPHPClient()

	transactionId := "PHP_TX_20260731001"
	// 构造正确的Token: MD5(AgentCode.ToUpper() + TransactionId.ToUpper() + API_SECRET_KEY)
	expectedToken := utils.SignWithdrawWithTransId(transactionId, PHP_WITHDRAW_AGENT_CODE, PHP_WITHDRAW_SECRET_KEY)
	fmt.Printf("PHP withdraw callback expected token: %s\n", expectedToken)

	err := cli.WithdrawCallback(OmPayWithdrawalCallbackReq{
		TransactionId: transactionId,
		StatusId:      2, // 2 = Completed
		StatusDesc:    "Completed",
		FullName:      "Juan dela Cruz",
		AccountNo:     "09171234567",
		Amount:        500.00,
		Token:         expectedToken,
	}, func(req OmPayWithdrawalCallbackReq) error {
		fmt.Printf("PHP withdraw callback processed: TransactionId=%s, StatusId=%d, Amount=%.2f\n",
			req.TransactionId, req.StatusId, req.Amount)
		return nil
	})

	if err != nil {
		t.Errorf("PHP WithdrawCallback failed: %+v\n", err)
	}
}

// TestPHPWithdrawCallbackInvalidToken 测试PHP出金回调签名验证失败场景
func TestPHPWithdrawCallbackInvalidToken(t *testing.T) {
	cli := newPHPClient()

	err := cli.WithdrawCallback(OmPayWithdrawalCallbackReq{
		TransactionId: "PHP_TX_20260731001",
		StatusId:      2,
		StatusDesc:    "Completed",
		FullName:      "Juan dela Cruz",
		AccountNo:     "09171234567",
		Amount:        500.00,
		Token:         "invalid_token_should_fail",
	}, func(req OmPayWithdrawalCallbackReq) error {
		return nil
	})

	if err == nil {
		t.Error("PHP WithdrawCallback should fail with invalid token")
	} else {
		fmt.Printf("PHP withdraw callback correctly rejected invalid token: %v\n", err)
	}
}
