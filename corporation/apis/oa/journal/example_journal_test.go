package journal_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/oa/journal"
)

func ExampleGetRecordList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := journal.GetRecordList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetRecordDetail() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := journal.GetRecordDetail(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetStatList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := journal.GetStatList(ctx, payload)

	fmt.Println(resp, err)
}
