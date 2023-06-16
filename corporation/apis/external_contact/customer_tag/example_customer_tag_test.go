package customer_tag_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/external_contact/customer_tag"
)

func ExampleGetCorpTagList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_tag.GetCorpTagList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddCorpTag() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_tag.AddCorpTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleEditCorpTag() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_tag.EditCorpTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelCorpTag() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_tag.DelCorpTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMarkTag() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := customer_tag.MarkTag(ctx, payload)

	fmt.Println(resp, err)
}
