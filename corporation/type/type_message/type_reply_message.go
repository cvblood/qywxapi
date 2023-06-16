package type_message

import "encoding/xml"

type CDATA string

func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

const (
	ReplyMsgTypeText  = "text"
	ReplyMsgTypeImage = "image"
	ReplyMsgTypeVoice = "voice"
	ReplyMsgTypeVideo = "video"
	ReplyMsgTypeMusic = "music"
	ReplyMsgTypeNews  = "news"
)

type ReplyMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   string
	MsgType      CDATA
}

/*
加密处理后 的 回复 消息体
<xml>
   <Encrypt><![CDATA[msg_encrypt]]></Encrypt>
   <MsgSignature><![CDATA[msg_signature]]></MsgSignature>
   <TimeStamp>timestamp</TimeStamp>
   <Nonce><![CDATA[nonce]]></Nonce>
</xml>

*/
type ReplyEncryptMessage struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string
	MsgSignature string
	TimeStamp    string
	Nonce        string
}

/*

<xml>
   <ToUserName><![CDATA[toUser]]></ToUserName>
   <FromUserName><![CDATA[fromUser]]></FromUserName>
   <CreateTime>1348831860</CreateTime>
   <MsgType><![CDATA[text]]></MsgType>
   <Content><![CDATA[this is a test]]></Content>
</xml>


*/
type ReplyMessageText struct {
	ReplyMessage
	Content CDATA
}

/*

<xml>
   <ToUserName><![CDATA[toUser]]></ToUserName>
   <FromUserName><![CDATA[fromUser]]></FromUserName>
   <CreateTime>1348831860</CreateTime>
   <MsgType><![CDATA[image]]></MsgType>
   <Image>
       <MediaId><![CDATA[media_id]]></MediaId>
   </Image>
</xml>


*/
type ReplyMessageImage struct {
	ReplyMessage
	Image struct {
		MediaId CDATA
	}
}

/*

<xml>
   <ToUserName><![CDATA[toUser]]></ToUserName>
   <FromUserName><![CDATA[fromUser]]></FromUserName>
   <CreateTime>1357290913</CreateTime>
   <MsgType><![CDATA[voice]]></MsgType>
   <Voice>
       <MediaId><![CDATA[media_id]]></MediaId>
   </Voice>
</xml>


*/
type ReplyMessageVoice struct {
	ReplyMessage

	Voice struct {
		MediaId CDATA
	}
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[video]]></MsgType>
  <Video>
    <MediaId><![CDATA[media_id]]></MediaId>
    <Title><![CDATA[title]]></Title>
    <Description><![CDATA[description]]></Description>
  </Video>
</xml>
*/
type ReplyMessageVideo struct {
	ReplyMessage
	Video struct {
		MediaId     CDATA
		Title       CDATA
		Description CDATA
	}
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[news]]></MsgType>
  <ArticleCount>1</ArticleCount>
  <Articles>
    <item>
      <Title><![CDATA[title1]]></Title>
      <Description><![CDATA[description1]]></Description>
      <PicUrl><![CDATA[picurl]]></PicUrl>
      <Url><![CDATA[url]]></Url>
    </item>
  </Articles>
</xml>

*/
type ReplyMessageNews struct {
	ReplyMessage
	ArticleCount string
	Articles     struct {
		Item []struct {
			Title       CDATA
			Description CDATA
			PicUrl      CDATA
			URL         CDATA
		}
	}
}
