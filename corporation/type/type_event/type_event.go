package type_event

import (
	"encoding/xml"
)

const (
	EventTypeTaskCardClick = "taskcard_click" // 任务卡片事件
)

type Event struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Event        string   `xml:"Event"`
}

/*
<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[FromUser]]></FromUserName>
<CreateTime>123456789</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[taskcard_click]]></Event>
<EventKey><![CDATA[key111]]></EventKey>
<TaskId><![CDATA[taskid111]]></TaskId >
<AgentId>1</AgentId>
</xml>
*/
type EventTaskCardClick struct {
	Event
	EventKey string `xml:"EventKey"`
	TaskId   string `xml:"TaskId"`
	AgentId  string `xml:"AgentId"`
}
