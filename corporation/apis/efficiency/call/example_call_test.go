package call_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/efficiency/call"
)

func ExampleGetDialRecord() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := call.GetDialRecord(ctx, payload)

	fmt.Println(resp, err)
}
