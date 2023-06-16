package type_event

const (
	EventTypeApproval = "open_approval_change" // 审批状态事件
)

/*
<xml>
 <ToUserName>wwddddccc7775555aaa</ToUserName>
  <FromUserName>sys</FromUserName>
  <CreateTime>1527838022</CreateTime>
  <MsgType>event</MsgType>
  <Event>open_approval_change</Event>
  <AgentID>1</AgentID>
  <ApprovalInfo>
    <ThirdNo>201806010001</ThirdNo>
    <OpenSpName>付款</OpenSpName>
    <OpenTemplateId>1234567890</OpenTemplateId>
    <OpenSpStatus>1</OpenSpStatus>
    <ApplyTime>1527837645</ApplyTime>
    <ApplyUserName>xiaoming</ApplyUserName>
    <ApplyUserId>1</ApplyUserId>
    <ApplyUserParty>产品部</ApplyUserParty>
    <ApplyUserImage>http://www.qq.com/xxx.png</ApplyUserImage>
    <ApprovalNodes>
      <ApprovalNode>
        <NodeStatus>1</NodeStatus>
        <NodeAttr>1</NodeAttr>
        <NodeType>1</NodeType>
        <Items>
          <Item>
            <ItemName>xiaohong</ItemName>
            <ItemUserId>2</ItemUserId>
            <ItemImage>http://www.qq.com/xxx.png</ItemImage>
            <ItemStatus>1</ItemStatus>
            <ItemSpeech></ItemSpeech>
            <ItemOpTime>0</ItemOpTime>
          </Item>
		  <Item>
            <ItemName>xiaohong</ItemName>
            <ItemUserId>2</ItemUserId>
            <ItemImage>http://www.qq.com/xxx.png</ItemImage>
            <ItemStatus>1</ItemStatus>
            <ItemSpeech></ItemSpeech>
            <ItemOpTime>0</ItemOpTime>
          </Item>
        </Items>
      </ApprovalNode>
    </ApprovalNodes>
    <NotifyNodes>
      <NotifyNode>
        <ItemName>xiaogang</ItemName>
        <ItemUserId>3</ItemUserId>
        <ItemImage>http://www.qq.com/xxx.png</ItemImage>
      </NotifyNode>
    </NotifyNodes>
    <approverstep>0</approverstep>
  </ApprovalInfo>
</xml>
*/
type EventApproval struct {
	Event
	AgentID      string `xml:"AgentID"`
	ApprovalInfo struct {
		Text           string `xml:",chardata"`
		ThirdNo        string `xml:"ThirdNo"`
		OpenSpName     string `xml:"OpenSpName"`
		OpenTemplateId string `xml:"OpenTemplateId"`
		OpenSpStatus   string `xml:"OpenSpStatus"`
		ApplyTime      string `xml:"ApplyTime"`
		ApplyUserName  string `xml:"ApplyUserName"`
		ApplyUserId    string `xml:"ApplyUserId"`
		ApplyUserParty string `xml:"ApplyUserParty"`
		ApplyUserImage string `xml:"ApplyUserImage"`
		ApprovalNodes  struct {
			Text         string `xml:",chardata"`
			ApprovalNode struct {
				Text       string `xml:",chardata"`
				NodeStatus string `xml:"NodeStatus"`
				NodeAttr   string `xml:"NodeAttr"`
				NodeType   string `xml:"NodeType"`
				Items      struct {
					Text string `xml:",chardata"`
					Item []struct {
						Text       string `xml:",chardata"`
						ItemName   string `xml:"ItemName"`
						ItemUserId string `xml:"ItemUserId"`
						ItemImage  string `xml:"ItemImage"`
						ItemStatus string `xml:"ItemStatus"`
						ItemSpeech string `xml:"ItemSpeech"`
						ItemOpTime string `xml:"ItemOpTime"`
					} `xml:"Item"`
				} `xml:"Items"`
			} `xml:"ApprovalNode"`
		} `xml:"ApprovalNodes"`
		NotifyNodes struct {
			Text       string `xml:",chardata"`
			NotifyNode struct {
				Text       string `xml:",chardata"`
				ItemName   string `xml:"ItemName"`
				ItemUserId string `xml:"ItemUserId"`
				ItemImage  string `xml:"ItemImage"`
			} `xml:"NotifyNode"`
		} `xml:"NotifyNodes"`
		Approverstep string `xml:"approverstep"`
	} `xml:"ApprovalInfo"`
}
