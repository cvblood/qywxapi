// Package custom_app OA/自建应用
package custom_app

import (
	"bytes"

	"github.com/cvblood/qywxapi/corporation"
)

const (
	apiGetOpenApprovalData = "/cgi-bin/corp/getopenapprovaldata"
)

/*
查询自建应用审批单当前状态

开发者也可主动查询审批单的当前审批状态。

See: https://work.weixin.qq.com/api/doc/90000/90135/90269

POST https://qyapi.weixin.qq.com/cgi-bin/corp/getopenapprovaldata?access_token=ACCESS_TOKEN
*/
func GetOpenApprovalData(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetOpenApprovalData, bytes.NewReader(payload), "application/json;charset=utf-8")
}
