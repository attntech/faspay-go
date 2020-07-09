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
	CancelledBy        string `json:"cancellation_type"`
	CancellationType   string `json:"cancellation_type"`
	CancellationAmount string `json:"cancellation_amount"`
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
	Response     string           `json:"response"`
	MerchantID   string           `json:"merchant_id"`
	Merchant     string           `json:"merchant"`
	Payments     *PaymentsKredivo `json:"payments"`
	ResponseCode string           `json:"response_code"`
	ResponseDesc string           `json:"response_desc"`
}

func SendListInquiryPaymentType(params *ListPaymentTypeKredivoRequest) (result *ListPaymentTypeKredivoResponse, err error) {
	params.MerchantID = FaspayConfig.MerchandID
	params.Merchant = FaspayConfig.MerchantName
	params.Request = RequestPaymentTypeKredivo
	sendRequest, err := SendPost(nil, params, getListPaymentTypeKredivoURL())
	if err != nil {
		return nil, err
	}

	result = &ListPaymentTypeKredivoResponse{
		Response:     sendRequest["response"].(string),
		Merchant:     sendRequest["merchant"].(string),
		MerchantID:   sendRequest["merchant_id"].(string),
		Payments:     sendRequest["payments"].(*PaymentsKredivo),
		ResponseDesc: sendRequest["response_desc"].(string),
		ResponseCode: sendRequest["response_code"].(string),
	}
	return
}

func SendCancelTransactionKredivo(params *CancelTransactionRequestKredivo) (result *CancelTransactionResponseKrevido, err error) {
	params.MerchantID = FaspayConfig.MerchandID
	params.Merchant = FaspayConfig.MerchantName
	params.Request = RequestPaymentTypeKredivo
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
