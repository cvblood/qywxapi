// Package health_qrcode 家校应用/复学码
package health_qrcode

import (
	"bytes"

	"github.com/cvblood/qywxapi/corporation"
)

const (
	apiGetTeacherCustomizeHealthInfo = "/cgi-bin/school/user/get_teacher_customize_health_info"
	apiGetStudentCustomizeHealthInfo = "/cgi-bin/school/user/get_student_customize_health_info"
	apiGetHealthQrcode               = "/cgi-bin/school/user/get_health_qrcode"
)

/*
获取老师健康信息



See: https://work.weixin.qq.com/api/doc/90000/90135/93744

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/get_teacher_customize_health_info?access_token=ACCESS_TOKEN
*/
func GetTeacherCustomizeHealthInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetTeacherCustomizeHealthInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取学生健康信息



See: https://work.weixin.qq.com/api/doc/90000/90135/93745

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/get_student_customize_health_info?access_token=ACCESS_TOKEN
*/
func GetStudentCustomizeHealthInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetStudentCustomizeHealthInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取师生健康码



See: https://work.weixin.qq.com/api/doc/90000/90135/93746

POST https://qyapi.weixin.qq.com/cgi-bin/school/user/get_health_qrcode?access_token=ACCESS_TOKEN
*/
func GetHealthQrcode(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetHealthQrcode, bytes.NewReader(payload), "application/json;charset=utf-8")
}
