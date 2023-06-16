package util_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/util"
)

func ExampleGetApiDomainIp() {
	var ctx *corporation.App

	resp, err := util.GetApiDomainIp(ctx)

	fmt.Println(resp, err)
}

func ExampleGetCallbackIp() {
	var ctx *corporation.App

	resp, err := util.GetCallbackIp(ctx)

	fmt.Println(resp, err)
}
