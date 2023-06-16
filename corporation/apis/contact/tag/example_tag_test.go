package tag_test

import (
	"fmt"
	"net/url"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/contact/tag"
)

func ExampleCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := tag.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := tag.Update(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := tag.Delete(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *corporation.App

	params := url.Values{}
	resp, err := tag.Get(ctx, params)

	fmt.Println(resp, err)
}

func ExampleAddTagUsers() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := tag.AddTagUsers(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelTagUsers() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := tag.DelTagUsers(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleList() {
	var ctx *corporation.App

	resp, err := tag.List(ctx)

	fmt.Println(resp, err)
}
