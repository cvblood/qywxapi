package customer_assign_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/external_contact/customer_assign"
)

func ExampleGetUnassignedList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_assign.GetUnassignedList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTransfer() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_assign.Transfer(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetTransferResult() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_assign.GetTransferResult(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupChatTransfer() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_assign.GroupChatTransfer(ctx, payload)

	fmt.Println(resp, err)
}
