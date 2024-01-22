package response

type Response struct {
	Data struct {
	} `json:"data" description:"数据"`
	Message string `json:"message" description:"消息"`
	Status  int    `json:"status" description:"状态"`
	TraceId string `json:"traceId" description:"追踪ID"`
}

type Message struct {
	Message string `json:"message" description:"消息"`
	Status  int    `json:"status" description:"状态"`
	TraceId string `json:"traceId" description:"追踪ID"`
	Data    struct {
		MessageId string `json:"messageId" description:"消息ID"`
	} `json:"data" description:"数据"`
}
