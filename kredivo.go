package faspay

type ItemKredivo struct {
	ID      string `json:"id"`
	Product string `json:"product"`
	Qty     string `json:"qty"`
	Amount  string `json:"amount"`
}

type CancelTransactionRequestKredivo struct {
	Request            string `json:"request"`
	MerchantID         string `json:"merchant_id"`
	Merchant           string `json:"merchant"`
	TrxID              string `json:"trx_id"`
	BillNo             string `json:"bill_no"`
	BillTotal          string `json:"bill_total"`
	CancelledBy        string `json:"cancelled_by"`
	CancellationType   string `json:"cancellation_type"`
	CancellationAmount string `json:"cancellation_amount"`
	CancellationReason string `json:"cancellation_reason"`
	Signature          string `json:"signature"`
}

type CancelTransactionResponseKrevido struct {
	Response     string `json:"response"`
	MerchantID   string `json:"merchant_id"`
	Merchant     string `json:"merchant"`
	TrxID        string `json:"trx_id"`
	BillNo       string `json:"bill_no"`
	FraudStatus  string `json:"fraud_status"`
	ResponseCode string `json:"response_code"`
	ResponseDesc string `json:"response_desc"`
}
type ListPaymentTypeKredivoRequest struct {
	Request    string         `json:"request"`
	MerchantID string         `json:"merchant_id"`
	Merchant   string         `json:"merchant"`
	BillTotal  string         `json:"bill_total"`
	Item       []*ItemKredivo `json:"item"`
	Signature  string         `json:"signature"`
}

type PaymentsKredivo struct {
	RawMonthlyInstallment         string `json:"raw_monthly_installment"`
	Name                          string `json:"name"`
	Amount                        string `json:"amount"`
	Installment                   string `json:"installment_amount"`
	Rate                          string `json:"raw_amount"`
	DownPayment                   string `json:"down_payment"`
	MonthlyInstallment            string `json:"monthly_installment"`
	DiscountentMonthlyInstallment string `json:"discounted_monthly_installment"`
	Tenure                        string `json:"tenure"`
	ID                            string `json:"id"`
}

type ListPaymentTypeKredivoResponse struct {
	Response     string      `json:"response"`
	MerchantID   string      `json:"merchant_id"`
	Merchant     string      `json:"merchant"`
	Payments     interface{} `json:"payments"`
	ResponseCode string      `json:"response_code"`
	ResponseDesc string      `json:"response_desc"`
}

func SendPaymentKredivo(params *PaymentDebitRequest) (result *PaymentResponse, err error) {

	// please refer to this doc https://faspay.co.id/docs/index-business.html#url-endpoint-post-data-3136

	params.MerchantId = FaspayConfig.MerchandID
	params.Request = RequestPaymentKredivo
	params.Signature = GetPaymentSignature(params.BillNo)
	params.BillGross = params.BillGross + "00"
	params.BillTotal = params.BillTotal + "00"
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

func SendListInquiryPaymentTypeKredivo(params *ListPaymentTypeKredivoRequest) (result *ListPaymentTypeKredivoResponse, err error) {
	params.MerchantID = FaspayConfig.MerchandID
	params.Merchant = FaspayConfig.MerChantUser
	params.Request = RequestPaymentTypeKredivo
	params.Signature = getSignatureKredivo()
	sendRequest, err := SendPost(nil, params, getListPaymentTypeKredivoURL())
	if err != nil {
		return nil, err
	}

	result = &ListPaymentTypeKredivoResponse{
		Response:     sendRequest["response"].(string),
		Merchant:     sendRequest["merchant"].(string),
		MerchantID:   sendRequest["merchant_id"].(string),
		Payments:     sendRequest["payments"].(interface{}),
		ResponseDesc: sendRequest["response_desc"].(string),
		ResponseCode: sendRequest["response_code"].(string),
	}
	return
}

func SendCancelTransactionKredivo(params *CancelTransactionRequestKredivo) (result *CancelTransactionResponseKrevido, err error) {
	params.MerchantID = FaspayConfig.MerchandID
	params.Merchant = FaspayConfig.MerChantUser
	params.Request = RequestCancelTransactionKredivo
	params.CancellationReason = "Tidak jadi bayar"
	params.CancellationType = "FULL"
	params.CancelledBy = FaspayConfig.MerChantUser
	params.Signature = GetPaymentSignature(params.BillNo)
	sendRequest, err := SendPost(nil, params, getCancelTransactionKredivoURL())
	if err != nil {
		return nil, err
	}

	result = &CancelTransactionResponseKrevido{
		Response:     sendRequest["response"].(string),
		Merchant:     sendRequest["merchant"].(string),
		MerchantID:   sendRequest["merchant_id"].(string),
		TrxID:        sendRequest["trx_id"].(string),
		BillNo:       sendRequest["bill_no"].(string),
		FraudStatus:  sendRequest["fraud_status"].(string),
		ResponseDesc: sendRequest["response_desc"].(string),
		ResponseCode: sendRequest["response_code"].(string),
	}
	return
}
