package app_test

import (
	"fmt"
	"net/url"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/app"
)

func ExampleAgentGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := app.AgentGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleAgentList() {
	var ctx *corporation.App

	resp, err := app.AgentList(ctx)

	fmt.Println(resp, err)
}

func ExampleAgentSet() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := app.AgentSet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMenuCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	params := url.Values{}
	resp, err := app.MenuCreate(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleMenuGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := app.MenuGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleMenuDelete() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := app.MenuDelete(ctx, params)

	fmt.Println(resp, err)
}

func ExampleSetWorkbenchTemplate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := app.SetWorkbenchTemplate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetWorkbenchTemplate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := app.GetWorkbenchTemplate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetWorkbenchData() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := app.SetWorkbenchData(ctx, payload)

	fmt.Println(resp, err)
}
