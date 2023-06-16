package invoice_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/invoice"
)

func ExampleGetInvoiceInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := invoice.GetInvoiceInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateInvoiceStatus() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := invoice.UpdateInvoiceStatus(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateStatusBatch() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := invoice.UpdateStatusBatch(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetInvoiceInfoBatch() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := invoice.GetInvoiceInfoBatch(ctx, payload)

	fmt.Println(resp, err)
}
