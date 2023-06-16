package msgaudit_test

import (
	"fmt"
	"net/url"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/msgaudit"
)

func ExampleGetRobotInfo() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := msgaudit.GetRobotInfo(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetPermitUserList() {
	var ctx *corporation.App

	resp, err := msgaudit.GetPermitUserList(ctx)

	fmt.Println(resp, err)
}

func ExampleCheckSingleAgree() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := msgaudit.CheckSingleAgree(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupchatGet() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := msgaudit.GroupchatGet(ctx, payload)

	fmt.Println(resp, err)
}
