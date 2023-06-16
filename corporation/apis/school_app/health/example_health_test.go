package health_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/school_app/health"
)

func ExampleGetHealthReportStat() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health.GetHealthReportStat(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetReportJobids() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health.GetReportJobids(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetReportJobInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health.GetReportJobInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetReportAnswer() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health.GetReportAnswer(ctx, payload)

	fmt.Println(resp, err)
}
