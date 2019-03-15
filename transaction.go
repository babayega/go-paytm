package paytm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PaytmResponse struct {
	TXNID       string `json:"TXNID"`
	BANKTXNID   string `json:"BANKTXNID"`
	ORDERID     string `json:"ORDERID"`
	TXNAMOUNT   string `json:"TXNAMOUNT"`
	STATUS      string `json:"STATUS"`
	TXNTYPE     string `json:"TXNTYPE"`
	GATEWAYNAME string `json:"GATEWAYNAME"`
	RESPCODE    string `json:"RESPCODE"`
	RESPMSG     string `json:"RESPMSG"`
	BANKNAME    string `json:"BANKNAME"`
	MID         string `json:"MID"`
	PAYMENTMODE string `json:"PAYMENTMODE"`
	REFUNDAMT   string `json:"REFUNDAMT"`
	TXNDATE     string `json:"TXNDATE"`
}

func GetTransactionStatus(pd map[string]string, MID string, url string) (pr PaytmResponse, err error) {
	var (
		req  *http.Request
		resp *http.Response
		//body []byte
	)

	jsonStr := fmt.Sprintf(`{"MID":"%s","ORDERID":"%s","CHECKSUMHASH":"%s"}`, MID, pd["ORDERID"], pd["CHECKSUMHASH"])

	req, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return pr, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return pr, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&pr)

	if err != nil {
		return pr, err
	}
	return pr, nil
}
