package material_test

import (
	"fmt"
	"net/url"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/material"
)

func ExampleUpload() {
	var ctx *corporation.App

	media := ""
	params := url.Values{}
	resp, err := material.Upload(ctx, media, params)

	fmt.Println(resp, err)
}

func ExampleUploadImg() {
	var ctx *corporation.App

	media := ""
	resp, err := material.UploadImg(ctx, media)

	fmt.Println(resp, err)
}
