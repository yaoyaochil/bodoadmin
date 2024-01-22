package mqtt_server

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/mqtt-server/config"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/mqtt-server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/mqtt-server/router"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/mqtt-server/utils"
)

var Mqtt = utils.UtilsApp.MqttClass

type MqttServerPlugin struct{}

func CreateMqttServerPlugin(host, port, topic string) *MqttServerPlugin {
	global.MqttGlobalConfig.Host = host
	global.MqttGlobalConfig.Port = port
	global.MqttGlobalConfig.Topic = topic

	err := Mqtt.InitMqttClient(config.Options{
		Host:  global.MqttGlobalConfig.Host,
		Port:  global.MqttGlobalConfig.Port,
		Topic: global.MqttGlobalConfig.Topic,
	}, "middle_server")
	if err != nil {
		fmt.Printf("Error initializing MQTT client: %v\n", err)
	}

	// 在单独的 goroutine 中运行 MQTT 监听
	//go startMqttListener()
	return &MqttServerPlugin{}
}

// 启动 MQTT 监听
func startMqttListener() {
	// 订阅主题，并设置回调函数
	if err := Mqtt.SubscribeToTopic(global.MqttGlobalConfig.Topic, handleMessage); err != nil {
		fmt.Printf("Error subscribing to topic: %v\n", err)
	}
}

// 处理接收到的消息
func handleMessage(client mqtt.Client, msg mqtt.Message) {
	// 处理接收到的消息
	fmt.Printf("Received message on topic %s: %s\n", msg.Topic(), msg.Payload())
}

func (*MqttServerPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitSwitchRouterGroup(group)
}

func (*MqttServerPlugin) RouterPath() string {
	return "mqtt"
}
