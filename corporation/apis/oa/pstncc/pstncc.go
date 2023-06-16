// Package pstncc OA/紧急通知应用
package pstncc

import (
	"bytes"

	"github.com/cvblood/qywxapi/corporation"
)

const (
	apiCall      = "/cgi-bin/pstncc/call"
	apiGetStates = "/cgi-bin/pstncc/getstates"
)

/*
发起语音电话

通过此接口发起语音电话，提醒员工查看应用推送的重要消息。

See: https://work.weixin.qq.com/api/doc/90000/90135/91627

POST https://qyapi.weixin.qq.com/cgi-bin/pstncc/call?access_token=ACCESS_TOKEN
*/
func Call(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCall, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取接听状态

通过此接口，了解员工是否已接听语音电话。

See: https://work.weixin.qq.com/api/doc/90000/90135/91628

POST https://qyapi.weixin.qq.com/cgi-bin/pstncc/getstates?access_token=ACCESS_TOKEN
*/
func GetStates(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetStates, bytes.NewReader(payload), "application/json;charset=utf-8")
}
