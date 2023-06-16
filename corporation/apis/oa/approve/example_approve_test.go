package approve_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/oa/approve"
)

func ExampleGetTemplateDetail() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.GetTemplateDetail(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleApplyEvent() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.ApplyEvent(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetApprovalInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.GetApprovalInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetApprovalDetail() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.GetApprovalDetail(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetApprovalData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.GetApprovalData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCorpConf() {
	var ctx *corporation.App

	resp, err := approve.GetCorpConf(ctx)

	fmt.Println(resp, err)
}

func ExampleGetUserVacationQuota() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.GetUserVacationQuota(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetOneUserQuota() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := approve.SetOneUserQuota(ctx, payload)

	fmt.Println(resp, err)
}
