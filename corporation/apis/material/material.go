// Package material 素材管理
package material

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/cvblood/qywxapi/corporation"
)

const (
	apiUpload    = "/cgi-bin/media/upload"
	apiUploadImg = "/cgi-bin/media/uploadimg"
	apiGet       = "/cgi-bin/media/get"
	apiJssdk     = "/cgi-bin/media/get/jssdk"
)

/*
上传临时素材



See: https://work.weixin.qq.com/api/doc/90000/90135/90253

POST(@media) https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
*/
func Upload(ctx *corporation.App, media string, params url.Values) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

	}()
	return ctx.Client.HTTPPost(apiUpload+"?"+params.Encode(), r, m.FormDataContentType())
}

/*
上传图片



See: https://work.weixin.qq.com/api/doc/90000/90135/90256

POST(@media) https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN
*/
func UploadImg(ctx *corporation.App, media string) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

	}()
	return ctx.Client.HTTPPost(apiUploadImg, r, m.FormDataContentType())
}

/*
获取临时素材



See: https://work.weixin.qq.com/api/doc/90000/90135/90254

GET https://qyapi.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
*/
func Get(ctx *corporation.App, params url.Values, header http.Header) (resp *http.Response, err error) {
	accessToken, err := ctx.AccessToken.GetAccessTokenHandler(ctx)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodGet, corporation.WXServerUrl+apiGet+"?access_token="+accessToken+"&"+params.Encode(), nil)
	if err != nil {
		return
	}

	req.Header = header

	return http.DefaultClient.Do(req)
}

/*
获取高清语音素材



See: https://work.weixin.qq.com/api/doc/90000/90135/90255

GET https://qyapi.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
*/
func Jssdk(ctx *corporation.App, params url.Values) (resp *http.Response, err error) {
	accessToken, err := ctx.AccessToken.GetAccessTokenHandler(ctx)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodGet, corporation.WXServerUrl+apiJssdk+"?access_token="+accessToken+"&"+params.Encode(), nil)
	if err != nil {
		return
	}

	return http.DefaultClient.Do(req)
}
