package health_qrcode_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/school_app/health_qrcode"
)

func ExampleGetTeacherCustomizeHealthInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health_qrcode.GetTeacherCustomizeHealthInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetStudentCustomizeHealthInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health_qrcode.GetStudentCustomizeHealthInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetHealthQrcode() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := health_qrcode.GetHealthQrcode(ctx, payload)

	fmt.Println(resp, err)
}
