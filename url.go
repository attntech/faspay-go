package faspay

func getPaymentChannelUrl() string {
	return FaspayConfig.Url + "cvr/100001/10"
}

func getPaymentUrl() string {
	return FaspayConfig.Url + "cvr/300011/10"
}

func getPaymentInquiryStatusUrl() string {
	return FaspayConfig.Url + "cvr/100004/10"
}

func getPaymentRedirectUrl(billno, trxID string) string {
	endpoint := "pws/100003/2830000010100000/"
	if FaspayConfig.Env != "prod" {
		endpoint = "pws/100003/0830000010100000/"
	}
	return FaspayConfig.Url + endpoint + GetPaymentSignature(billno) + "?trx_id=" + trxID + "&merchant_id=" + FaspayConfig.MerchandID + "&bill_no=" + billno
}

func getPaymenDebitCancelUrl() string {
	return FaspayConfig.Url + "cvr/100005/10"
}
