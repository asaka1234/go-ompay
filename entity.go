package go_ompay

// ----------pre order-------------------------

type OMPayDepositReq struct {
	SerialNo          string `json:"serialNo" mapstructure:"serialNo"`                   //商户的单号  //1
	Currency          string `json:"currency" mapstructure:"currency"`                   //1
	Amount            string `json:"amount" mapstructure:"amount"`                       //Return URL after the payment is done.
	ClientAccountName string `json:"ClientAccountName" mapstructure:"ClientAccountName"` //Client's Registered Full name in account (KYC)
	//这几个让sdk来搞
	//MerchantCode string `json:"merchantCode"` //商户id
	//ReturnUrl    string `json:"returnUrl"`    //前端回跳地址
	//NotifyUrl    string `json:"notifyUrl"`    //非必填,回调通知接口
	//Token        string `json:"token"`        //签名MD5(serialNo + {API KEY} + {Secret KEY} + amount)
}

// ------------------------------------------------------------

type OmPayDepositCallbackReq struct {
	Info         string `json:"Info" mapstructure:"Info"` // Payment information. (E.g. Approved, Rejected)
	MerchantCode string `json:"MerchantCode" mapstructure:"MerchantCode"`
	SerialNo     string `json:"SerialNo" mapstructure:"SerialNo"`
	CurrencyCode string `json:"CurrencyCode" mapstructure:"CurrencyCode"`
	Amount       string `json:"Amount" mapstructure:"Amount"`
	Status       int    `json:"Status" mapstructure:"Status"` //0 = Waiting for payment,1 = Payment Approved ,2 = Rejected
	Token        string `json:"Token" mapstructure:"Token"`   //MD5(serialNo + {API KEY} + {Secret KEY})
}
