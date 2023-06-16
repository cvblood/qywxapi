package message_test

import (
	"fmt"
	"net/url"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/message"
)

func ExampleSend() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.Send(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateTaskcard() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.UpdateTaskcard(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAppchatCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.AppchatCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAppchatUpdate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.AppchatUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAppchatGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := message.AppchatGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleAppchatSend() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.AppchatSend(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleLinkedcorpMessageSend() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.LinkedcorpMessageSend(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetStatistics() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := message.GetStatistics(ctx, payload)

	fmt.Println(resp, err)
}
