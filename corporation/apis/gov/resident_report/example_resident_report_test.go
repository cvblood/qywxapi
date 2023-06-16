package resident_report_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/gov/resident_report"
)

func ExampleGetGridInfo() {
	var ctx *corporation.App

	resp, err := resident_report.GetGridInfo(ctx)

	fmt.Println(resp, err)
}

func ExampleGetCorpStatus() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := resident_report.GetCorpStatus(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserStatus() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := resident_report.GetUserStatus(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCategoryStatistic() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := resident_report.CategoryStatistic(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetOrderList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := resident_report.GetOrderList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetOrderInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := resident_report.GetOrderInfo(ctx, payload)

	fmt.Println(resp, err)
}
