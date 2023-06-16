// Package call 效率工具/公费电话
package call

import (
	"bytes"

	"github.com/cvblood/qywxapi/corporation"
)

const (
	apiGetDialRecord = "/cgi-bin/dial/get_dial_record"
)

/*
获取公费电话拨打记录

企业可通过此接口，按时间范围拉取成功接通的公费电话拨打记录。

See: https://work.weixin.qq.com/api/doc/90000/90135/93662

POST https://qyapi.weixin.qq.com/cgi-bin/dial/get_dial_record?access_token=ACCESS_TOKEN
*/
func GetDialRecord(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetDialRecord, bytes.NewReader(payload), "application/json;charset=utf-8")
}
