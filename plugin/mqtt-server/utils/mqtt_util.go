package utils

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/mqtt-server/config"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/mqtt-server/global"
	"time"
)

type MqttClass struct {
	client mqtt.Client
}

func (m *MqttClass) InitMqttClient(config config.Options, clientID string) error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", config.Host, config.Port))
	opts.SetClientID(clientID)

	// 设置连接断开时的回调函数
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v\n", err)
		m.tryReconnect(opts)
	})

	// 初始化 MQTT 客户端
	global.MqttGlobalConfig.MqttCommon = mqtt.NewClient(opts)
	if token := global.MqttGlobalConfig.MqttCommon.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("MQTT connection error: %v", token.Error())
	}

	fmt.Println("Connected to MQTT broker successfully!")

	// 启动心跳检测
	go m.startHeartbeat()
	return nil
}

func (m *MqttClass) tryReconnect(opts *mqtt.ClientOptions) {
	for !global.MqttGlobalConfig.MqttCommon.IsConnected() {
		fmt.Println("Attempting to reconnect...")
		if token := global.MqttGlobalConfig.MqttCommon.Connect(); token.Wait() && token.Error() != nil {
			fmt.Printf("Reconnection error: %v. Retrying...\n", token.Error())
			time.Sleep(5 * time.Second) // 5秒后尝试重连
		} else {
			fmt.Println("Reconnected successfully!")
			break
		}
	}
}

func (m *MqttClass) startHeartbeat() {
	for {
		if global.MqttGlobalConfig.MqttCommon.IsConnected() {
			// 发送心跳消息
			// TODO: 添加发送心跳消息的逻辑
		}
		time.Sleep(5 * time.Second) // 每5秒发送一次心跳
	}
}

// PublishMessage 发布消息到指定主题
func (m *MqttClass) PublishMessage(topic string, payload string) error {
	if token := global.MqttGlobalConfig.MqttCommon.Publish(topic, 1, false, payload); token.Wait() && token.Error() != nil {
		return fmt.Errorf("MQTT publish error: %v", token.Error())
	}
	return nil
}

// SubscribeToTopic 订阅指定主题的消息
func (m *MqttClass) SubscribeToTopic(topic string, callback func(client mqtt.Client, msg mqtt.Message)) error {
	if token := global.MqttGlobalConfig.MqttCommon.Subscribe(topic, 1, callback); token.Wait() && token.Error() != nil {
		return fmt.Errorf("MQTT subscribe error: %v", token.Error())
	}
	return nil
}
