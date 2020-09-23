package faspay

import (
	"os"
)

var FaspayConfig Config

type Config struct {
	MerchantName     string
	MerchandID       string
	MerChantUser     string
	MerchantPassword string
	Env              string
	Url              string
	TxnPassword      string
}

func Initialize() {

	FaspayConfig.MerchantName = os.Getenv("MERCHANT_NAME")
	if FaspayConfig.MerchantName == "" {
		panic("MERCHANT_NAME is empty. do this; export MERCHANT_NAME=XXXXX")
	}

	FaspayConfig.MerchandID = os.Getenv("MERCHANT_ID")
	if FaspayConfig.MerchandID == "" {
		panic("MERCHANT_ID is empty. do this; export MERCHANT_ID=XXXXX")
	}

	FaspayConfig.MerChantUser = os.Getenv("MERCHANT_USER")
	if FaspayConfig.MerChantUser == "" {
		panic("MERCHANT_USER is empty. do this; export MERCHANT_USER=XXXXX")
	}

	FaspayConfig.MerchantPassword = os.Getenv("MERCHANT_PASSWORD")
	if FaspayConfig.MerchantPassword == "" {
		panic("MERCHANT_PASSWORD is empty. do this; export MERCHANT_PASSWORD=XXXXX")
	}

	FaspayConfig.MerchantPassword = os.Getenv("MERCHANT_PASSWORD")
	if FaspayConfig.MerchantPassword == "" {
		panic("MERCHANT_PASSWORD is empty. do this; export MERCHANT_PASSWORD=XXXXX")
	}

	FaspayConfig.Env = os.Getenv("ENV")
	if FaspayConfig.Env == "" {
		panic("ENV is empty. do this; export ENV=XXXXX")
	}

	if FaspayConfig.Env != "prod" {
		FaspayConfig.Url = "https://dev.faspay.co.id/"
	} else {
		FaspayConfig.Url = "https://web.faspay.co.id/"
	}

}

func NewClient() {
	if FaspayConfig.Env != "prod" {
		FaspayConfig.Url = "https://dev.faspay.co.id/"
	} else {
		FaspayConfig.Url = "https://web.faspay.co.id/"
	}
}
