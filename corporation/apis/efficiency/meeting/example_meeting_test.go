package meeting_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/efficiency/meeting"
)

func ExampleCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting.Update(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCancel() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting.Cancel(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserMeetingId() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting.GetUserMeetingId(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting.GetInfo(ctx, payload)

	fmt.Println(resp, err)
}
