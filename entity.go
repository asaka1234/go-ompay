package go_ompay

type OMPayInitParams struct {
	MerchantInfo `yaml:",inline" mapstructure:",squash"`

	//银行转账地址
	DepositUrl       string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl"  yaml:"depositUrl"`
	DepositBackUrl   string `json:"depositBackUrl" mapstructure:"depositBackUrl" config:"depositBackUrl"  yaml:"depositBackUrl"`
	DepositFeBackUrl string `json:"depositFeBackUrl" mapstructure:"depositFeBackUrl" config:"depositFeBackUrl"  yaml:"depositFeBackUrl"`

	//fpx/qrcode支付
	FPXDepositUrl string `json:"fpxDepositUrl" mapstructure:"fpxDepositUrl" config:"fpxDepositUrl"  yaml:"fpxDepositUrl"`

	WithdrawUrl     string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`
	WithdrawBackUrl string `json:"withdrawBackUrl" mapstructure:"withdrawBackUrl" config:"withdrawBackUrl"  yaml:"withdrawBackUrl"`
}

type MerchantInfo struct {
	MerchantId        string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"` //貌似只是给deposit用的
	DepositApiKey     string `json:"depositApiKey" mapstructure:"depositApiKey" config:"depositApiKey"  yaml:"depositApiKey"`
	DepositSecretKey  string `json:"depositSecretKey" mapstructure:"depositSecretKey" config:"depositSecretKey"  yaml:"depositSecretKey"`
	WithdrawAgentCode string `json:"withdrawAgentCode" mapstructure:"withdrawAgentCode" config:"withdrawAgentCode"  yaml:"withdrawAgentCode"`
	WithdrawSecretKey string `json:"withdrawSecretKey" mapstructure:"withdrawSecretKey" config:"withdrawSecretKey"  yaml:"withdrawSecretKey"`
}

// ----------pre order-------------------------
//没有psp三方的订单号

type OMPayDepositReq struct {
	SerialNo          string `json:"serialNo" mapstructure:"serialNo"`                   //商户的唯一单号
	Currency          string `json:"currency" mapstructure:"currency"`                   //1
	Amount            string `json:"amount" mapstructure:"amount"`                       //Return URL after the payment is done.
	ClientAccountName string `json:"ClientAccountName" mapstructure:"ClientAccountName"` //Client's Registered Full name in account (KYC)
	//这几个让sdk来搞
	//MerchantCode string `json:"merchantCode"` //商户id
	//ReturnUrl    string `json:"returnUrl"`    //前端回跳地址
	//NotifyUrl    string `json:"notifyUrl"`    //非必填,回调通知接口
	//Token        string `json:"token"`        //签名MD5(serialNo + {API KEY} + {Secret KEY} + amount)
}

type OMPayFPXDepositReq struct {
	SerialNo string  `json:"serialNo" mapstructure:"serialNo"` //商户的唯一单号
	Currency string  `json:"currency" mapstructure:"currency"` //1
	Amount   string  `json:"amount" mapstructure:"amount"`     //Return URL after the payment is done.
	PayType  *string `json:"payType" mapstructure:"payType"`   //Currency IDR Only. (IDR Default 1003), - 1003 (VA), - 1004 (QRIS)
	//这几个让sdk来搞
	//MerchantCode string `json:"merchantCode"` //商户id
	//ReturnUrl    string `json:"returnUrl"`    //前端回跳地址
	//NotifyUrl    string `json:"notifyUrl"`    //非必填,回调通知接口
	//Token        string `json:"token"`        //签名MD5(serialNo + {API KEY} + {Secret KEY} + amount)
}

type OMPayFPXDepositResponse struct {
	Success      bool   `json:"success" mapstructure:"success"`           //成功与否
	Data         string `json:"data" mapstructure:"data"`                 //Return payment URL if success.
	MerchantCode string `json:"merchantCode" mapstructure:"merchantCode"` //Unique ID for each merchant.
	SerialNo     string `json:"serialNo" mapstructure:"serialNo"`         //Unique Transaction Id from request
	Message      string `json:"message" mapstructure:"message"`           //Error message if not success.
}

// ------------------------------------------------------------

// psp是发了一个post请求过来, 如果处理好了就返回 success ,失败就返回fail
// 样式: https://{your notify or callback url}?Info=Approved&MerchantCode=60-00000100-0000123&SerialNo=AA0123456&CurrencyCode=MYR&Amount=500.00&Status=1&Token=***
type OmPayDepositCallbackReq struct {
	Info         string  `json:"Info" mapstructure:"Info"` // Payment information. (E.g. Approved, Rejected)
	MerchantCode string  `json:"MerchantCode" mapstructure:"MerchantCode"`
	SerialNo     string  `json:"SerialNo" mapstructure:"SerialNo"` //商户的唯一单号
	CurrencyCode string  `json:"CurrencyCode" mapstructure:"CurrencyCode"`
	Amount       float64 `json:"Amount" mapstructure:"Amount"`
	Status       int     `json:"Status" mapstructure:"Status"` //0 = Waiting for payment,1 = Payment Approved ,2 = Rejected
	Token        string  `json:"Token" mapstructure:"Token"`   //MD5(serialNo + {API KEY} + {Secret KEY})
}

//=============================提现====================================

type OMPayWithdrawalReq struct {
	UserRef           string  `json:"UserRef" mapstructure:"UserRef"`                                         //感觉是唯一,可以是随机生成. Your reference Id for this request
	TransactionId     string  `json:"TransactionId" mapstructure:"TransactionId"`                             //商户的订单唯一id
	FullName          string  `json:"FullName,omitempty" mapstructure:"FullName,omitempty"`                   //Full name of the user of the bank. (Required except USDT)
	AccountNo         string  `json:"AccountNo" mapstructure:"AccountNo"`                                     //Bank Account Number / USDT Address
	BankCode          string  `json:"BankCode" mapstructure:"BankCode"`                                       //Bank Name / USDT Types (TRX, ETH)
	BankRegister      string  `json:"BankRegister,omitempty" mapstructure:"BankRegister,omitempty"`           //Bank Branch
	BankRegisterState string  `json:"BankRegisterState,omitempty" mapstructure:"BankRegisterState,omitempty"` //Bank Branch State
	BankRegisterCity  string  `json:"BankRegisterCity,omitempty" mapstructure:"BankRegisterCity,omitempty"`   //Bank Branch City
	Amount            float64 `json:"Amount" mapstructure:"Amount"`
	Currency          string  `json:"Currency" mapstructure:"Currency"` //币种
	//让sdk设置
	//AgentCode   string `json:"AgentCode" mapstructure:"AgentCode"`                         //给merchant分配的
	//CallbackURL string `json:"CallbackURL,omitempty" mapstructure:"CallbackURL,omitempty"` //回调地址
	//Token       string `json:"Token" mapstructure:"Token"`                                 //签名, MD5({Agent Code} + UserRef.ToUpper() + {Secret KEY})
}

type OMPayWithdrawalResp struct {
	Success      bool   `json:"success" mapstructure:"success"`
	UserRef      string `json:"userRef" mapstructure:"userRef"`     //唯一id, Your reference Id for this request.
	ResultRef    int    `json:"resultRef" mapstructure:"resultRef"` //Our result reference id
	ErrorCode    int    `json:"errorCode" mapstructure:"errorCode"` //0 是正确，非0是错误
	ErrorMessage string `json:"errorMessage" mapstructure:"errorMessage"`
	//错误时的返回
	HasError bool   `json:"HasError" mapstructure:"HasError"`
	Info     string `json:"Info" mapstructure:"Info"`
}

// psp是发了一个post请求过来, 如果处理好了就返回 success ,失败就返回fail
// 发送的url样式： https://{your_callback_url}/?Token=***&TransactionId=ABC123456&StatusDesc=Complete
type OmPayWithdrawalCallbackReq struct {
	TransactionId string  `json:"TransactionId" mapstructure:"TransactionId"` //商户的唯一单号
	StatusId      int     `json:"StatusId" mapstructure:"StatusId"`           //2 = Completed, 3 = Rejected
	StatusDesc    string  `json:"StatusDesc" mapstructure:"StatusDesc"`       //提现状态: Completed, Rejected
	FullName      string  `json:"FullName" mapstructure:"FullName"`           //Full name of the user of the bank
	AccountNo     string  `json:"AccountNo" mapstructure:"AccountNo"`         //Bank Account Number
	Amount        float64 `json:"Amount" mapstructure:"Amount"`
	Token         string  `json:"Token" mapstructure:"Token"` // 签名 MD5(AgentCode.ToUpper() + TransactionId.ToUpper() + API_SECRET_KEY)
}
