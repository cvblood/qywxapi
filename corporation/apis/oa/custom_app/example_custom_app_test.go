package custom_app_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/oa/custom_app"
)

func ExampleGetOpenApprovalData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := custom_app.GetOpenApprovalData(ctx, payload)

	fmt.Println(resp, err)
}
