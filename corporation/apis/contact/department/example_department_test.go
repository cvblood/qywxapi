package department_test

import (
	"fmt"
	"net/url"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/contact/department"
)

func ExampleCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := department.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := department.Update(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := department.Delete(ctx, params)

	fmt.Println(resp, err)
}

func ExampleList() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := department.List(ctx, params)

	fmt.Println(resp, err)
}
