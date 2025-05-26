package go_ompay

const (
	MERCHANT_ID = "60-00000555-98958489" //商户id
	API_KEY     = "B0328E1E-55C5-4BB5-BB32-5DEB819934CE"
	SECRET_KEY  = "74AE81ABF1444C93B28F286F58046552"

	WITHDRAW_AGENT_CODE = "46CE1F81-ACAA-49A8-B950-177322FA05E1"
	WITHDRAW_SECRET_KEY = "857F13A98DAC453F83D1616AA1355A98"

	//--------

	DEPOSIT_URL  = "https://api.doitwallet.asia/Merchant/Pay"        //充值
	WITHDRAW_URL = "https://api.doitwallet.asia/api/wallet/withdraw" //提现 payout

	DEPOSIT_CALLBACK_URL    = "http://127.0.0.1/order/post"
	DEPOSIT_FE_CALLBACK_URL = "http://127.0.0.1/order/post"

	WITHDRAW_CALLBACK_URL = "http://127.0.0.1/order/post"
)
