package school_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/school"
)

func ExampleGetSubscribeQrCode() {
	var ctx *corporation.App

	resp, err := school.GetSubscribeQrCode(ctx)

	fmt.Println(resp, err)
}

func ExampleSetSubscribeMode() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := school.SetSubscribeMode(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetSubscribeMode() {
	var ctx *corporation.App

	resp, err := school.GetSubscribeMode(ctx)

	fmt.Println(resp, err)
}

func ExampleSend() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := school.Send(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleConvertToOpenid() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := school.ConvertToOpenid(ctx, payload)

	fmt.Println(resp, err)
}
