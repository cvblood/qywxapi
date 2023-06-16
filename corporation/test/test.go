// Package test 模拟微信服务器 测试
package test

import (
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/cvblood/qywxapi/corporation"
)

var MockCorporation *corporation.Corporation
var MockApp *corporation.App
var MockSvr *httptest.Server
var MockSvrHandler *http.ServeMux
var onceSetup sync.Once

// 初始化测试环境
func Setup() {
	onceSetup.Do(func() {

		MockCorporation = corporation.New(corporation.Config{
			Corpid: "Cropid",
		})
		MockApp = MockCorporation.NewApp(corporation.AppConfig{
			AgentId:        "AgentId",
			Secret:         "SECRET",
			Token:          "TOKEN",
			EncodingAESKey: "EncodingAESKey",
		})

		// Mock Server
		MockSvrHandler = http.NewServeMux()
		MockSvr = httptest.NewServer(MockSvrHandler)
		corporation.WXServerUrl = MockSvr.URL // 拦截发往微信服务器的请求

		// Mock access token
		MockSvrHandler.HandleFunc("/cgi-bin/gettoken", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"access_token":"ACCESS_TOKEN","expires_in":7200}`))
		})
	})
}
