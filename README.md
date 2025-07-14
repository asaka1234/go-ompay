整体
==============
1. 充值: https://api.doitwallet.asia/Documents/DepositAPI.pdf
2. 提现: https://api.doitwallet.asia/Documents/PayoutAPI.pdf
3. 二维码: https://api.doitwallet.asia/Documents/FPXAPI.pdf
4. 支持currency: MYR, IDR, SGD, THB, VND



鉴权
==============
1. query的md5签名


充值
==============
1. 后端帮签名,生成一个收银台url.  随后fe打开这个url即可进行支付


提现
===============
1. withdraw是一个post请求,application/x-www-form-urlencoded类型
2. 可以接口中动态定义callback url