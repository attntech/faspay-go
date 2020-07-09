package faspay

const (
	LinkAja               = "302"
	XLTunai               = "303"
	BRIMOCash             = "400"
	BRIEPay               = "401"
	PermataNet            = "402"
	BCAKlikPay            = "405"
	MAYBANKVA             = "408"
	CIMBClicks            = "700"
	DanamonOnlineBanking  = "701"
	VABCA                 = "702"
	Sakuku                = "704"
	IndomaretPaymentPoint = "706"
	AlfaGroup             = "707"
	DanamonVA             = "708"
	Kredivo               = "709"
	BRIVA                 = "800"
	BNIVA                 = "801"
	MandiriVA             = "802"
	Akulaku               = "807"
	OVO                   = "812"
	Maybank2U             = "814"

	PaymentStatusCodeBelumDiProses = iota
	PaymentStatusCodeDalamProses
	PaymentStatusCodePaymentSukses
	PaymentStatusCodePaymentGagal
	PaymentStatusCodePaymentReversal
	PaymentStatusCodeTagihanTidakDitemukan
	PaymentStatusCodePaymentExpired
	PaymentStatusCodePaymentCancelled
	PaymentStatusCodeUnknown

	TerminalOriginWeb           = "10"
	TerminalOriginMobAppIOS     = "22"
	TerminalOriginMobAppAndroid = "21"

	PayTypeFullSettlement = "1"
	PayTypeInstallment    = "2"
	PayTypeMixed          = "3"

	PaymentPlanFullSettlement = "1"
	PaymentPlanInstallment    = "2"

	TenorFullPayment = "00"
	Tenor3Months     = "03"
	Tenor6Months     = "06"
	Tenor12Months    = "12"

	RequestPaymentDebit       = "Transmisi Info Detil Pembelian"
	RequestPaymentChannel     = "Daftar Payment Channel"
	RequestPaymentDebitCancel = "Canceling Payment"
	RequestPaymentInquiry     = "Pengecekan Status Pembayaran"
	RequestPaymentTypeKredivo = "Pengecekan Payment Type Kredivo"

	ResponseCodeSuccess                 = "00"
	ResponseCodeFailure                 = "01"
	ResponseCodeCancelPaymentByCustomer = "17"
)
