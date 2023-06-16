package living_test

import (
	"fmt"
	"net/url"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/school_app/living"
)

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

func ExampleGetUnwatchStat() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := living.GetUnwatchStat(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDeleteReplayData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := living.DeleteReplayData(ctx, payload)

	fmt.Println(resp, err)
}
