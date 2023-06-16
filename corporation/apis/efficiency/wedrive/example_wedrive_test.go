package wedrive_test

import (
	"fmt"

	"github.com/cvblood/qywxapi/corporation"
	"github.com/cvblood/qywxapi/corporation/apis/efficiency/wedrive"
)

func ExampleSpaceCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.SpaceCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSpaceRename() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.SpaceRename(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSpaceDismiss() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.SpaceDismiss(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSpaceInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.SpaceInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSpaceAclAdd() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.SpaceAclAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSpaceAclDel() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.SpaceAclDel(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSpaceSetting() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.SpaceSetting(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSpaceShare() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.SpaceShare(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleFileList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.FileList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleFileUpload() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.FileUpload(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleFileDownload() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.FileDownload(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleFileCreate() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.FileCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleFileRename() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.FileRename(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleFileMove() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.FileMove(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleFileDelete() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.FileDelete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleFileInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.FileInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleFileAclAdd() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.FileAclAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleFileAclDel() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.FileAclDel(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleFileSetting() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.FileSetting(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleFileShare() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := wedrive.FileShare(ctx, payload)

	fmt.Println(resp, err)
}
