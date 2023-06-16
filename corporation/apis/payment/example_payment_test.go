package payment_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/payment"
)

func ExampleAddMerchant() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := payment.AddMerchant(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMerchant() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := payment.GetMerchant(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelMerchant() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := payment.DelMerchant(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetMchUseScope() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := payment.SetMchUseScope(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetBillList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := payment.GetBillList(ctx, payload)

	fmt.Println(resp, err)
}
