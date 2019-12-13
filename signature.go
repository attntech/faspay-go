package faspay

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func GetPaymentChannelSignature() string {
	return getSha1(getMD5Hash(FaspayConfig.MerChantUser + FaspayConfig.MerchantPassword))
}

func GetPaymentSignature(billNo string) string {
	return getSha1(getMD5Hash(FaspayConfig.MerChantUser + FaspayConfig.MerchantPassword + billNo))
}

func GetPaymentNotificationSignature(billNo, paymentStatusCode string) string {
	return getSha1(getMD5Hash(FaspayConfig.MerChantUser + FaspayConfig.MerchantPassword + billNo + paymentStatusCode))
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func getSha1(str string) (signature string) {
	h := sha1.New()
	h.Write([]byte(str))
	signature = hex.EncodeToString(h.Sum(nil))
	return
}
