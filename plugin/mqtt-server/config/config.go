package config

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Options struct {
	Host       string
	Port       string
	Topic      string
	MqttCommon mqtt.Client
}
