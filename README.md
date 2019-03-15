# go-paytm
Checksum utilities for PayTM written in golang

## Primary Functions
This library provides three primary functions

* Generate checksum
* Verify checksum
* Get Transaction status


## Configuration
Before integrating you need to set up the following configuration

```golang
#config.go

const (
	PaytmMerchantKey = `xxxxxxxxxxxxxxxx`
	MID              = `xxxxxxxxxxxxxxxxxxxx`
	INDUSTRY_TYPE_ID = `Retail`
	CHANNEL_ID       = `WAP`
	WEBSITE          = `APPSTAGING`
	CALLBACK_URL     = `https://securegw-stage.paytm.in/theia/paytmCallback?ORDER_ID=`
	TransactionStatusAPI = `https://securegw-stage.paytm.in/merchant-status/getTxnStatus`
)
```

## Examples
* GetChecksumFromArray  

```golang
func GetChecksumFromArray() {
	var (
		orderId  = "order456"
		customId = "custom001"
		amount   = "1.00"
	)
	paramList := map[string]string{
		"MID":              MID,
		"INDUSTRY_TYPE_ID": INDUSTRY_TYPE_ID,
		"CHANNEL_ID":       CHANNEL_ID,
		"WEBSITE":          WEBSITE,
		"CALLBACK_URL":     CALLBACK_URL + orderID,
		"ORDER_ID":         orderID,
		"CUST_ID":          customID,
		"TXN_AMOUNT":       amount,
	}

	checksum, err := paytm.GetChecksumFromArray(paramList, PaytmMerchantKey)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("checksum = ", checksum)
	}
}
```

* VerifyCheckum  

```golang
func VerifyCheckum() {
	result := `{"ORDERID":"27000364888888", "MID":"Pay85623985963121", "TXNID":"20180710111212800110168868500018912", "TXNAMOUNT":"1.00", "PAYMENTMODE":"DC", "CURRENCY":"INR", "TXNDATE":"2018-07-10 15:04:56.0", "STATUS":"TXN_SUCCESS", "RESPCODE":"01", "RESPMSG":"Txn Success", "GATEWAYNAME":"HDFC", "BANKTXNID":"4036217121962950", "BANKNAME":"HDFC Bank", "CHECKSUMHASH":"TabnoADfqfWjI3twGIsjTRb97iDXlJSjq3S+fWOOtsz608mo+6JsAy600VZR/uimKR/46bdjrwgREQh4uF0L6IBeuhAhabyzUfJ5s2i5wps="}`
	resultList := map[string]string{}
	json.Unmarshal([]byte(result), &resultList)

	fmt.Println("resultList = ", resultList)

	ok := paytm.VerifyCheckum(resultList, PaytmMerchantKey)
	fmt.Println("ok = ", ok)
}
```

* GetTransactionStatus  

```golang
func GetTransactionStatus()  {
		paramList := map[string]string{
		"ORDERID":              "order456",
		"CHECKSUMHASH": 		"very long and secure hash",
	}
	res, err := paytm.GetTransactionStatus(paramList, MID, TransactionStatusAPI)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return
	}
	fmt.Println("res = ", res)
}
```
