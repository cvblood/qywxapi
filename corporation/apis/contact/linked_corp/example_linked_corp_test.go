package linked_corp_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/contact/linked_corp"
)

func ExampleGetPermList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := linked_corp.GetPermList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := linked_corp.Get(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSimpleList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := linked_corp.SimpleList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUserList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := linked_corp.UserList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDepartmentList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := linked_corp.DepartmentList(ctx, payload)

	fmt.Println(resp, err)
}
