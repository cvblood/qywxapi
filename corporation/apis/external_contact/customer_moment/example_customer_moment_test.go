package customer_moment_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/external_contact/customer_moment"
)

func ExampleGetMomentList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_moment.GetMomentList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMomentTask() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_moment.GetMomentTask(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMomentCustomerList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_moment.GetMomentCustomerList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMomentSendResult() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_moment.GetMomentSendResult(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMomentComments() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_moment.GetMomentComments(ctx, payload)

	fmt.Println(resp, err)
}
