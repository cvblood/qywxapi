package pstncc_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/oa/pstncc"
)

func ExampleCall() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := pstncc.Call(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetStates() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := pstncc.GetStates(ctx, payload)

	fmt.Println(resp, err)
}
