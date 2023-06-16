package async_batch_test

import (
	"fmt"
	"net/url"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/contact/async_batch"
)

func ExampleUser() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := async_batch.User(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReplaceUser() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := async_batch.ReplaceUser(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReplaceParty() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := async_batch.ReplaceParty(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetResult() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := async_batch.GetResult(ctx, params)

	fmt.Println(resp, err)
}
