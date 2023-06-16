package patrol_report_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/gov/patrol_report"
)

func ExampleGetGridInfo() {
	var ctx *corporation.App

	resp, err := patrol_report.GetGridInfo(ctx)

	fmt.Println(resp, err)
}

func ExampleGetCorpStatus() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := patrol_report.GetCorpStatus(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserStatus() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := patrol_report.GetUserStatus(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCategoryStatistic() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := patrol_report.CategoryStatistic(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetOrderList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := patrol_report.GetOrderList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetOrderInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := patrol_report.GetOrderInfo(ctx, payload)

	fmt.Println(resp, err)
}
