package model

// TextMessage 文字消息
type TextMessage struct {
	Type int `json:"type" description:"消息类型"`
	Data struct {
		EventId   string `json:"eventId" description:"事件id"`
		EventBody struct {
			IslandSourceId string `json:"islandSourceId" description:"来源群id"`
			MessageType    int    `json:"messageType" description:"消息类型，1：文字消息，2：图片消息，3：视频消息，4：分享消息，5：文件消息，6：卡片消息，7：红包消息"`
			MessageBody    struct {
				Content string `json:"content" description:"消息内容"`
			} `json:"messageBody" description:"消息体"`
			DodoSourceId string `json:"dodoSourceId" description:"发送者用户ID"`
			Member       struct {
				JoinTime string `json:"joinTime" description:"加群时间"`
				NickName string `json:"nickName" description:"群昵称"`
			} `json:"member" description:"成员信息"`
			MessageId string `json:"messageId" description:"消息ID"`
			Personal  struct {
				AvatarUrl string `json:"avatarUrl" description:"DoDo头像"`
				NickName  string `json:"nickName" description:"DoDo昵称"`
				Sex       int    `json:"sex" description:"性别，-1：保密，0：女，1：男"`
			} `json:"personal" description:"个人信息"`
			ChannelId string `json:"channelId" description:"来源频道ID"`
		} `json:"eventBody" description:"事件内容"`
		EventType string `json:"eventType" description:"事件类型"`
		Timestamp int64  `json:"timestamp" description:"时间戳"`
	} `json:"data"`
	Version string `json:"version" description:"业务版本"`
}

type CallBackType struct {
	Type int `json:"type"`
	Data struct {
		EventId   string `json:"eventId"`
		EventType string `json:"eventType"`
		Timestamp int64  `json:"timestamp"`
	} `json:"data"`
	Version string `json:"version"`
}

type PersonalMessageEvent struct {
	Type int `json:"type"`
	Data struct {
		EventBody struct {
			IslandSourceId string `json:"islandSourceId"`
			DodoSourceId   string `json:"dodoSourceId"`
			Personal       struct {
				NickName  string `json:"nickName"`
				AvatarUrl string `json:"avatarUrl"`
				Sex       int    `json:"sex"`
			} `json:"personal"`
			MessageId   string `json:"messageId"`
			MessageType int    `json:"messageType"`
			MessageBody struct {
				Content string `json:"content"`
			} `json:"messageBody"`
		} `json:"eventBody"`
		EventId   string `json:"eventId"`
		EventType string `json:"eventType"`
		Timestamp int64  `json:"timestamp"`
	} `json:"data"`
	Version string `json:"version"`
}
