package customer_group_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/external_contact/customer_group"
)

func ExampleList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_group.List(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_group.Get(ctx, payload)

	fmt.Println(resp, err)
}
