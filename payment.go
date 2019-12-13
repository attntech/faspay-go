package faspay

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

	params.MerchantId = FaspayConfig.MerchandID
	params.Request = RequestPaymentDebit
	params.Signature = GetPaymentSignature(params.BillNo)
	sendRequest, err := SendPost(nil, params, getPaymentUrl())
	if err != nil {
		return nil, err
	}

	result = &PaymentResponse{
		Response:     sendRequest["response"].(string),
		TrxID:        sendRequest["trx_id"].(string),
		ResponseCode: sendRequest["response_code"].(string),
		ResponseDesc: sendRequest["response_desc"].(string),
		BillNo:       sendRequest["bill_no"].(string),
		BillItems:    sendRequest["bill_items"].([]interface{}),
		RedirectURL:  getPaymentRedirectUrl(params.BillNo, sendRequest["trx_id"].(string)),
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
