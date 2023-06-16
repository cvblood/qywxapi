package meeting_room_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/oa/meeting_room"
)

func ExampleAdd() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.Add(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.List(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleEdit() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.Edit(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDel() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.Del(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetBookingInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.GetBookingInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBook() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.Book(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCancelBook() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.CancelBook(ctx, payload)

	fmt.Println(resp, err)
}
