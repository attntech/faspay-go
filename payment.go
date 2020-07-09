package faspay

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
)

type PaymentNotification struct {
	Request           string `json:"request"`
	TrxID             string `json:"trx_id"`
	MerchantID        string `json:"merchant_id"`
	Merchant          string `json:"merchant"`
	BillNo            string `json:"bill_no"`
	PaymentReff       string `json:"payment_reff"`
	PaymentDate       string `json:"payment_date"`
	PaymentStatusCode string `json:"payment_status_code"`
	PaymentStatusDesc string `json:"payment_status_desc"`
	BillTotal         string `json:"bill_total"`
	PaymentTotal      string `json:"payment_total"`
	PaymentChannelUID string `json:"payment_channel_uid"`
	PaymentChannel    string `json:"payment_channel"`
	Signature         string `json:"signature"`
}

type PaymentNotificationResponse struct {
	Response     string `json:"response"`
	TrxID        string `json:"trx_id"`
	MerchantID   string `json:"merchant_id"`
	Merchant     string `json:"merchant"`
	BillNo       string `json:"bill_no"`
	ResponseCode string `json:"response_code"`
	ResponseDesc string `json:"response_desc"`
	ResponseDate string `json:"response_date"`
}

type Item struct {
	Product     string `json:"product"`
	Qty         string `json:"qty"`
	Amount      string `json:"amount"`
	PaymentPlan string `json:"payment_plan"`
	MerchantID  string `json:"merchant_id"`
	Tenor       string `json:"tenor"`
}

type PaymentResponse struct {
	Response     string        `json:"response"`
	TrxID        string        `json:"trx_id"`
	MerchantID   string        `json:"merchant_id"`
	Merchant     string        `json:"merchant"`
	BillNo       string        `json:"bill_no"`
	BillItems    []interface{} `json:"bill_items"`
	ResponseCode string        `json:"response_code"`
	ResponseDesc string        `json:"response_desc"`
	RedirectURL  string        `json:"redirect_url"`
}

type PaymentDebitRequest struct {
	Request                    string  `json:"request"`
	MerchantId                 string  `json:"merchant_id"`
	BillNo                     string  `json:"bill_no"`
	BillReff                   string  `json:"bill_reff"`
	BillDate                   string  `json:"bill_date"`
	BillExpired                string  `json:"bill_expired"`
	BillDesc                   string  `json:"bill_desc"`
	BillCurrency               string  `json:"bill_currency"`
	BillGross                  string  `json:"bill_gross"`
	BillMiscfee                string  `json:"bill_miscfee"`
	BillTotal                  string  `json:"bill_total"`
	CustNo                     string  `json:"cust_no"`
	CustName                   string  `json:"cust_name"`
	PaymentChannel             string  `json:"payment_channel"`
	PayType                    string  `json:"pay_type"`
	BankUserid                 string  `json:"bank_userid"`
	Msisdn                     string  `json:"msisdn"`
	Email                      string  `json:"email"`
	Terminal                   string  `json:"terminal"`
	BillingName                string  `json:"billing_name"`
	BillingLastname            string  `json:"billing_lastname"`
	BillingAddress             string  `json:"billing_address"`
	BillingAddressCity         string  `json:"billing_address_city"`
	BillingAddressRegion       string  `json:"billing_address_region"`
	BillingAddressState        string  `json:"billing_address_state"`
	BillingAddressPoscode      string  `json:"billing_address_poscode"`
	BillingMsisdn              string  `json:"billing_msisdn"`
	BillingAddressCountryCode  string  `json:"billing_address_country_code"`
	ReceiverNameForShipping    string  `json:"receiver_name_for_shipping"`
	ShippingLastname           string  `json:"shipping_lastname"`
	ShippingAddress            string  `json:"shipping_address"`
	ShippingAddressCity        string  `json:"shipping_address_city"`
	ShippingAddressRegion      string  `json:"shipping_address_region"`
	ShippingAddressState       string  `json:"shipping_address_state"`
	ShippingAddressPoscode     string  `json:"shipping_address_poscode"`
	ShippingMsisdn             string  `json:"shipping_msisdn"`
	ShippingAddressCountryCode string  `json:"shipping_address_country_code"`
	Item                       []*Item `json:"item"`
	Reserve1                   string  `json:"reserve1"`
	Reserve2                   string  `json:"reserve2"`
	Signature                  string  `json:"signature"`
}

type PaymentWithCardRequest struct {
	PaymentMethod              string
	TxnPassword                string
	MerchantTranid             string
	Currencycode               string
	Amount                     string
	Custname                   string
	Custemail                  string
	Description                string
	ReturnURL                  string
	Signature                  string
	BillingAddress             string
	BillingAddressCity         string
	BillingAddressRegion       string
	BillingAddressState        string
	BillingAddressPoscode      string
	BillingAddressCountryCode  string
	ReceiverNameForShipping    string
	ShippingAddress            string
	ShippingAddressCity        string
	ShippingAddressRegion      string
	ShippingAddressState       string
	ShippingAddressPoscode     string
	ShippingAddressCountryCode string
	Shippingcost               string
	PhoneNo                    string
	Mref1                      string
	Mref2                      string
	Mref3                      string
	Mref4                      string
	Mref5                      string
	Mref6                      string
	Mref7                      string
	Mref8                      string
	Mref9                      string
	Mref10                     string
	Mparam1                    string
	Mparam2                    string
	CustomerRef                string
	PymtInd                    string
	PymtCriteria               string
	PymtToken                  string
	Frisk1                     string
	Frisk2                     string
	DomicileAddress            string
	DomicileAddressCity        string
	DomicileAddressRegion      string
	DomicileAddressState       string
	DomicileAddressPoscode     string
	DomicileAddressCountryCode string
	DomicilePhoneNo            string
	HandshakeUrl               string
	HandshakeParam             string
	StyleMerchantName          string
	StyleOrderSummary          string
	StyleOrderNo               string
	StyleOrderDesc             string
	StyleAmount                string
	StyleBackgroundLeft        string
	StyleButtonCancel          string
	StyleFontCancel            string
}

type PaymentChannelRequest struct {
	Request    string `json:"request"`
	MerchandID string `json:"merchant_id"`
	Merchant   string `json:"merchant"`
	Signature  string `json:"signature"`
}

type PaymentChannelResponse struct {
	Response       string        `json:"response"`
	MerchandID     string        `json:"merchant_id"`
	Merchant       string        `json:"merchant"`
	PaymentChannel []interface{} `json:"payment_channel"`
}

type PaymentDebitCancelRequest struct {
	Request       string `json:"request"`
	TrxID         string `json:"trx_id"`
	MerchandID    string `json:"merchant_id"`
	Merchant      string `json:"merchant"`
	BillNo        string `json:"bill_no"`
	PaymentCancel string `json:"payment_cancel"`
	Signature     string `json:"signature"`
}

type PaymentDebitCancelReponse struct {
	Response          string `json:"response"`
	TrxID             string `json:"trx_id"`
	MerchandID        string `json:"merchant_id"`
	Merchant          string `json:"merchant"`
	BillNo            string `json:"bill_no"`
	TrxStatusCode     string `json:"trx_status_code"`
	TrxStatusDesc     string `json:"trx_status_desc"`
	PaymentStatusCode string `json:"payment_status_code"`
	PaymentStatusDesc string `json:"payment_status_desc"`
	PaymentCancel     string `json:"payment_cancel"`
	ResponseCode      string `json:"response_code"`
	ResponseDesc      string `json:"response_desc"`
}

type PaymentInquiryRequest struct {
	Request    string `json:"request"`
	TrxID      string `json:"payment_channel"`
	MerchandID string `json:"merchant_id"`
	BillNo     string `json:"bill_no"`
	Signature  string `json:"signature"`
}

type PaymentInquiryResponse struct {
	Response          string `json:"response"`
	TrxID             string `json:"payment_channel"`
	MerchandID        string `json:"merchant_id"`
	Merchant          string `json:"merchant"`
	BillNo            string `json:"bill_no"`
	PaymentReff       string `json:"payment_reff"`
	PaymentDate       string `json:"payment_date"`
	PaymentStatusCode string `json:"payment_status_code"`
	PaymentStatusDesc string `json:"payment_status_desc"`
	ResponseCode      string `json:"response_code"`
	ResponseDesc      string `json:"response_desc"`
}

type VaStaticResponse struct {
	Response     string `json:"response"`
	VaNumber     string `json:"va_number"`
	Amount       string `json:"amount"`
	CustName     string `json:"cust_name"`
	ResponseCode string `json:"response_code"`
}

func GetPaymentChannel() (result *PaymentChannelResponse, err error) {

	payload := &PaymentChannelRequest{
		Request:    RequestPaymentChannel,
		MerchandID: FaspayConfig.MerchandID,
		Merchant:   FaspayConfig.MerchantName,
		Signature:  GetPaymentChannelSignature(),
	}

	sendRequest, err := SendPost(nil, payload, getPaymentChannelUrl())
	if err != nil {
		return nil, err
	}

	result = &PaymentChannelResponse{
		Response:       sendRequest["response"].(string),
		MerchandID:     sendRequest["merchant_id"].(string),
		Merchant:       sendRequest["merchant"].(string),
		PaymentChannel: sendRequest["payment_channel"].([]interface{}),
	}
	return
}

func SendPaymentDebit(params *PaymentDebitRequest) (result *PaymentResponse, err error) {

	// please refer to this doc https://faspay.co.id/docs/index-business.html#request-parameter-post-data
	params.MerchantId = FaspayConfig.MerchandID
	params.Request = RequestPaymentDebit
	params.Signature = GetPaymentSignature(params.BillNo)
	sendRequest, err := SendPost(nil, params, getPaymentUrl())
	if err != nil {
		return nil, err
	}

	if sendRequest != nil && sendRequest["response_error"] != nil {
		var responseCode map[string]interface{}
		responseCode = sendRequest["response_error"].(map[string]interface{})

		result = &PaymentResponse{
			ResponseDesc: responseCode["response_desc"].(string),
			ResponseCode: ResponseCodeFailure,
		}
		return result, nil
	}

	result = &PaymentResponse{
		Response:     sendRequest["response"].(string),
		TrxID:        sendRequest["trx_id"].(string),
		ResponseCode: sendRequest["response_code"].(string),
		ResponseDesc: sendRequest["response_desc"].(string),
		BillNo:       sendRequest["bill_no"].(string),
		BillItems:    sendRequest["bill_items"].([]interface{}),
		RedirectURL:  sendRequest["redirect_url"].(string),
	}
	return
}

func CancelPaymentDebit(reason, billNo, trxID string) (result *PaymentDebitCancelReponse, err error) {

	payload := &PaymentDebitCancelRequest{
		Request:       RequestPaymentChannel,
		MerchandID:    FaspayConfig.MerchandID,
		Merchant:      FaspayConfig.MerchantName,
		BillNo:        billNo,
		PaymentCancel: reason,
		TrxID:         trxID,
		Signature:     GetPaymentSignature(billNo),
	}

	sendRequest, err := SendPost(nil, payload, getPaymenDebitCancelUrl())
	if err != nil {
		return nil, err
	}

	result = &PaymentDebitCancelReponse{
		Response:          sendRequest["response"].(string),
		TrxID:             sendRequest["trx_id"].(string),
		ResponseCode:      sendRequest["response_code"].(string),
		ResponseDesc:      sendRequest["response_desc"].(string),
		BillNo:            sendRequest["bill_no"].(string),
		PaymentStatusCode: sendRequest["payment_status_code"].(string),
		PaymentStatusDesc: sendRequest["payment_status_desc"].(string),
		PaymentCancel:     sendRequest["payment_cancel"].(string),
	}
	return

}

func PaymentInquiryStatus(trxID, billNo string) (result *PaymentInquiryResponse, err error) {

	payload := &PaymentInquiryRequest{
		Request:    RequestPaymentChannel,
		MerchandID: FaspayConfig.MerchandID,
		BillNo:     billNo,
		TrxID:      trxID,
		Signature:  GetPaymentSignature(billNo),
	}

	sendRequest, err := SendPost(nil, payload, getPaymentInquiryStatusUrl())
	if err != nil {
		return nil, err
	}

	result = &PaymentInquiryResponse{
		Response:          sendRequest["response"].(string),
		TrxID:             sendRequest["trx_id"].(string),
		ResponseCode:      sendRequest["response_code"].(string),
		ResponseDesc:      sendRequest["response_desc"].(string),
		BillNo:            sendRequest["bill_no"].(string),
		PaymentStatusCode: sendRequest["payment_status_code"].(string),
		PaymentStatusDesc: sendRequest["payment_status_desc"].(string),
		PaymentDate:       sendRequest["payment_date"].(string),
		PaymentReff:       sendRequest["payment_reff"].(string),
	}
	return

}

func SendPaymentWithCard(params *PaymentWithCardRequest) (redirectURL string, err error) {

	signature := GetPaymentWithCardSignature(params.MerchantTranid, params.Amount)
	postData := url.Values{}
	postData.Set("TRANSACTIONTYPE", "1")
	postData.Set("RESPONSE_TYPE", "2")
	postData.Set("LANG", "")
	postData.Set("MERCHANTID", FaspayConfig.MerchandID)
	postData.Set("PAYMENT_METHOD", "1")
	postData.Set("TXN_PASSWORD", FaspayConfig.TxnPassword)
	postData.Set("MERCHANT_TRANID", params.MerchantTranid)
	postData.Set("CURRENCYCODE", "IDR")
	postData.Set("AMOUNT", params.Amount)
	postData.Set("CUSTNAME", params.Custname)
	postData.Set("CUSTEMAIL", params.Custemail)
	postData.Set("DESCRIPTION", params.Description)
	postData.Set("RETURN_URL", params.ReturnURL)
	postData.Set("SIGNATURE", signature)
	postData.Set("BILLING_ADDRESS", params.BillingAddress)
	postData.Set("BILLING_ADDRESS_CITY", params.BillingAddressCity)
	postData.Set("BILLING_ADDRESS_REGION", params.BillingAddressRegion)
	postData.Set("BILLING_ADDRESS_STATE", params.BillingAddressState)
	postData.Set("BILLING_ADDRESS_POSCODE", params.BillingAddressPoscode)
	postData.Set("BILLING_ADDRESS_COUNTRY_CODE", params.BillingAddressCountryCode)
	postData.Set("RECEIVER_NAME_FOR_SHIPPING", params.ReceiverNameForShipping)
	postData.Set("SHIPPING_ADDRESS", params.ShippingAddress)
	postData.Set("SHIPPING_ADDRESS_CITY", params.BillingAddressCity)
	postData.Set("SHIPPING_ADDRESS_REGION", params.ShippingAddressRegion)
	postData.Set("SHIPPING_ADDRESS_STATE", params.ShippingAddressState)
	postData.Set("SHIPPING_ADDRESS_POSCODE", params.ShippingAddressPoscode)
	postData.Set("SHIPPING_ADDRESS_COUNTRY_CODE", params.ShippingAddressCountryCode)
	postData.Set("SHIPPINGCOST", params.Shippingcost)
	postData.Set("PHONE_NO", params.PhoneNo)
	postData.Set("MREF1", params.Mref1)
	postData.Set("MREF2", params.Mref2)
	postData.Set("MREF3", params.Mref3)
	postData.Set("MREF4", params.Mref4)
	postData.Set("MREF5", params.Mref5)
	postData.Set("MREF6", params.Mref6)
	postData.Set("MREF7", params.Mref7)
	postData.Set("MREF8", params.Mref8)
	postData.Set("MREF9", params.Mref9)
	postData.Set("MREF10", params.Mref10)
	postData.Set("MPARAM1", params.Mparam1)
	postData.Set("MPARAM2", params.Mparam2)
	postData.Set("CUSTOMER_REF", params.CustomerRef)
	postData.Set("PYMT_IND", params.PymtInd)
	postData.Set("PYMT_CRITERIA", params.PymtCriteria)
	postData.Set("PYMT_TOKEN", params.PymtToken)
	postData.Set("FRISK1", params.Frisk1)
	postData.Set("FRISK2", params.Frisk2)
	postData.Set("DOMICILE_ADDRESS", params.DomicileAddress)
	postData.Set("DOMICILE_ADDRESS_CITY", params.DomicileAddressCity)
	postData.Set("DOMICILE_ADDRESS_REGION", params.DomicileAddressRegion)
	postData.Set("DOMICILE_ADDRESS_STATE", params.DomicileAddressState)
	postData.Set("DOMICILE_ADDRESS_POSCODE", params.DomicileAddressPoscode)
	postData.Set("DOMICILE_ADDRESS_COUNTRY_CODE", params.BillingAddressCountryCode)
	postData.Set("DOMICILE_PHONE_NO", params.PhoneNo)
	postData.Set("handshake_url", params.HandshakeUrl)
	postData.Set("handshake_param", "")
	postData.Set("style_merchant_name", "black")
	postData.Set("style_order_summary", "black")
	postData.Set("style_order_no", "black")
	postData.Set("style_order_desc", "black")
	postData.Set("style_amount", "black")
	postData.Set("style_background_left", "#fff")
	postData.Set("style_button_cancel", "grey")
	postData.Set("style_font_cancel", "white")

	request, err := http.NewRequest("POST", "https://fpgdev.faspay.co.id/payment", strings.NewReader(postData.Encode()))
	if err != nil {
		return
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Cookie", "PHPSESSID")

	// client := new(http.Client)
	// client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
	// 	return errors.New("Redirect")
	// }

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// response, err := http.DefaultClient.Do(request)
	// if err != nil {
	// 	return c.ReturnError(err)
	// }

	// client := new(http.Client)
	// client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
	// 	return errors.New("Redirect")
	// }

	response, err := client.Do(request)
	if err != nil {
		beego.Notice("RENE")
		return c.ReturnError(err)
	}

	defer response.Body.Close()
	fmt.Println(response.Request.URL.String())
	fmt.Println(response.StatusCode)

	url := ""
	if response.StatusCode == http.StatusFound { //status code 302
		beego.Notice("RENE BRO")

		url, err := response.Location()
		if err != nil {
			return c.ReturnError(err)
		}

		beego.Notice("BRORBO", url.String())
	}

}
