package faspay

func PaymentDebitExample() {
	item := &Item{
		Product:     "name of product",
		Qty:         "1",
		Amount:      "10000",
		PaymentPlan: PaymentPlanFullSettlement,
		MerchantID:  FaspayConfig.MerchandID,
		Tenor:       TenorFullPayment,
	}

	items := []*Item{item}

	payload := &PaymentDebitRequest{
		Request:                    "Transmisi Info Detil Pembelian",
		BillNo:                     "ORDERTEST10001",
		BillReff:                   "12345678",
		BillDate:                   "2019-12-10 20:48:10",
		BillExpired:                "2019-12-12 20:48:10",
		BillDesc:                   "Payment of ORDERTEST10001",
		BillCurrency:               "IDR",
		BillGross:                  "0",
		BillMiscfee:                "0",
		BillTotal:                  "10000",
		CustNo:                     "123",
		CustName:                   "Test Customer",
		PaymentChannel:             BRIVA,
		PayType:                    "1",
		BankUserid:                 "",
		Msisdn:                     "082313121",
		Email:                      "test@test.com",
		Terminal:                   TerminalOriginWeb,
		BillingName:                "Billing name",
		BillingLastname:            "Billing Lastname",
		BillingAddress:             "jalan cobaja",
		BillingAddressCity:         "Jakarta",
		BillingAddressRegion:       "DKI Jakarta",
		BillingAddressState:        "Indonesia",
		BillingAddressPoscode:      "12630",
		BillingMsisdn:              "082313131",
		BillingAddressCountryCode:  "ID",
		ReceiverNameForShipping:    "Test",
		ShippingLastname:           "Shipping Lastname",
		ShippingAddress:            "jalan cobaja",
		ShippingAddressCity:        "Jakarta",
		ShippingAddressRegion:      "DKI Jakarta",
		ShippingAddressState:       "Indonesia",
		ShippingAddressPoscode:     "12630",
		ShippingMsisdn:             "082313131",
		ShippingAddressCountryCode: "ID",
		Item:                       items,
		Signature:                  GetPaymentSignature("ORDERTEST10001"),
	}
	SendPaymentDebit(payload)
}
