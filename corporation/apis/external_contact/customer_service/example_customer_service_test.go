package customer_service_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/external_contact/customer_service"
)

func ExampleGetFollowUserList() {
	var ctx *corporation.App

	resp, err := customer_service.GetFollowUserList(ctx)

	fmt.Println(resp, err)
}

func ExampleAddContactWay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_service.AddContactWay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetContactWay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_service.GetContactWay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateContactWay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_service.UpdateContactWay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelContactWay() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_service.DelContactWay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCloseTempChat() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_service.CloseTempChat(ctx, payload)

	fmt.Println(resp, err)
}
