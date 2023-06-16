package customer_test

import (
	"fmt"
	"net/url"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/external_contact/customer"
)

func ExampleList() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := customer.List(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := customer.Get(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetByUser() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer.GetByUser(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleRemark() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer.Remark(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMobileHashcode() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer.GetMobileHashcode(ctx, payload)

	fmt.Println(resp, err)
}
