package corp_group_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/corp_group"
)

func ExampleListAppShareInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := corp_group.ListAppShareInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetToken() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := corp_group.GetToken(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTransferSession() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := corp_group.TransferSession(ctx, payload)

	fmt.Println(resp, err)
}
