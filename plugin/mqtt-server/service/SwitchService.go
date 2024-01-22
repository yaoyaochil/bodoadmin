package service

import "github.com/yaoyaochil/bodo-admin-server/server/plugin/mqtt-server/utils"

type SwitchService struct {
}

var Mqtt = utils.UtilsApp.MqttClass

func (s *SwitchService) SwitchOn() (err error) {
	err = Mqtt.PublishMessage("switch", "0")
	return err
}

func (s *SwitchService) SwitchOff() (err error) {
	err = Mqtt.PublishMessage("switch", "1")
	return err
}
