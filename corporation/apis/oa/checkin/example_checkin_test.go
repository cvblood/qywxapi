package checkin_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/oa/checkin"
)

func ExampleGetCorpCheckinOption() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.GetCorpCheckinOption(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCheckinOption() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.GetCheckinOption(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCheckinData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.GetCheckinData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCheckinDayData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.GetCheckinDayData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCheckinMonthData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.GetCheckinMonthData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCheckinScheduleList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.GetCheckinScheduleList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetCheckinScheduleList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.SetCheckinScheduleList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddCheckinUserFace() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := checkin.AddCheckinUserFace(ctx, payload)

	fmt.Println(resp, err)
}
