package customer_message_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/external_contact/customer_message"
)

func ExampleAddMsgTemplate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.AddMsgTemplate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGroupmsgList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GetGroupmsgList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGroupmsgTask() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GetGroupmsgTask(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGroupmsgSendResult() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GetGroupmsgSendResult(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSendWelcomeMsg() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.SendWelcomeMsg(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupWelcomeTemplateAdd() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GroupWelcomeTemplateAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupWelcomeTemplateEdit() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GroupWelcomeTemplateEdit(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupWelcomeTemplateGet() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GroupWelcomeTemplateGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGroupWelcomeTemplateDel() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_message.GroupWelcomeTemplateDel(ctx, payload)

	fmt.Println(resp, err)
}
