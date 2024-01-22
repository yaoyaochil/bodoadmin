package config

type Dodo struct {
	ClientID       string       // dodo机器人id
	Token          string       // dodo机器人Token
	Secret         string       // dodo机器人Secret
	Authorization  string       // dodo机器人Authorization
	IslandSourceId string       // dodo机器人群id
	WssURl         string       // dodo机器人wss地址
	InitFunc       func([]byte) // dodo机器人回调消息处理函数
}
