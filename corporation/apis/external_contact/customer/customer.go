// Package customer 客户联系/客户管理
package customer

import (
	"bytes"
	"net/url"

	"github.com/cvblood/qywxapi/corporation"
)

const (
	apiList              = "/cgi-bin/externalcontact/list"
	apiGet               = "/cgi-bin/externalcontact/get"
	apiGetByUser         = "/cgi-bin/externalcontact/batch/get_by_user"
	apiRemark            = "/cgi-bin/externalcontact/remark"
	apiGetMobileHashcode = "/cgi-bin/user/get_mobile_hashcode"
)

/*
获取客户列表

企业可通过此接口获取指定成员添加的客户列表。客户是指配置了客户联系功能的成员所添加的外部联系人。没有配置客户联系功能的成员，所添加的外部联系人将不会作为客户返回。

See: https://work.weixin.qq.com/api/doc/90000/90135/92113

GET https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list?access_token=ACCESS_TOKEN&userid=USERID
*/
func List(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiList + "?" + params.Encode())
}

/*
获取客户详情

企业可通过此接口，根据外部联系人的userid（如何获取?），拉取客户详情。

See: https://work.weixin.qq.com/api/doc/90000/90135/92114

GET https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get?access_token=ACCESS_TOKEN&external_userid=EXTERNAL_USERID
*/
func Get(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGet + "?" + params.Encode())
}

/*
批量获取客户详情

企业/第三方可通过此接口获取指定成员添加的客户信息列表。

See: https://work.weixin.qq.com/api/doc/90000/90135/92994

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user?access_token=ACCESS_TOKEN
*/
func GetByUser(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetByUser, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改客户备注信息

企业可通过此接口修改指定用户添加的客户的备注信息。

See: https://work.weixin.qq.com/api/doc/90000/90135/92115

POST https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark?access_token=ACCESS_TOKEN
*/
func Remark(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiRemark, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取手机号随机串

支持企业获取手机号随机串，该随机串可直接在企业微信终端搜索手机号对应的微信用户。

See: https://work.weixin.qq.com/api/doc/90000/90135/92790

POST https://qyapi.weixin.qq.com/cgi-bin/user/get_mobile_hashcode?access_token=ACCESS_TOKEN
*/
func GetMobileHashcode(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetMobileHashcode, bytes.NewReader(payload), "application/json;charset=utf-8")
}
