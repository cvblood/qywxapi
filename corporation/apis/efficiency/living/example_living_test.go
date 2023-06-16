package living_test

import (
	"fmt"
	"net/url"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/efficiency/living"
)

func ExampleCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := living.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleModify() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := living.Modify(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCancel() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := living.Cancel(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDeleteReplayData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := living.DeleteReplayData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetLivingCode() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := living.GetLivingCode(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserAllLivingId() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := living.GetUserAllLivingId(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetLivingInfo() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := living.GetLivingInfo(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetWatchStat() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := living.GetWatchStat(ctx, payload)

	fmt.Println(resp, err)
}
