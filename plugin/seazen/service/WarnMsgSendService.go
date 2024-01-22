package service

import (
	"errors"
	"fmt"
	BODO_Global "github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/model"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/passport"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/util"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/service"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/wx_request"
	"strconv"
)

type WarnMsgSendService struct{}

// SendWarningTemplateMsg 发送销采预警消息
func (s *WarnMsgSendService) SendWarningTemplateMsg() (err error) {
	var sendLog model.WarnMsgInfoList
	if err = BODO_Global.BODO_DB.Model(&sendLog).Where("send_date = ?", util.GetToday()).First(&sendLog).Error; err == nil {
		return errors.New("今日已推送过,无需再推送")
	}

	shopEntitySaleInfoData, err := passport.PassportGroupApp.GagService.GetShopEntitySaleInfo()
	if err != nil {
		return
	}
	var haveDiffShop string
	var haveDiffShopCount int
	for _, v := range shopEntitySaleInfoData.Data.Data {
		if v.RecommendClaimDiffAbs != 0 {
			haveDiffShopCount++
			haveDiffShop += v.EntityName + ","
		}
	}
	shopEntityCheckData, err := passport.PassportGroupApp.GagService.GetShopEntityCheck()
	if err != nil {
		return
	}
	var notDeviceShop string
	var notDeviceShopCount int
	for _, v := range shopEntityCheckData.Data.Data {
		if v.IsTerminal != 0 {
			notDeviceShopCount++
			notDeviceShop += v.EntityName + ","
		}
	}
	if notDeviceShopCount == 0 && haveDiffShopCount == 0 {
		return errors.New("无需推送")
	}
	var OpendIDs []string
	if err = BODO_Global.BODO_DB.Model(model.SendMsgUserList{}).Where("is_enabled = ?", true).Pluck("open_id", &OpendIDs).Error; err != nil {
		return
	}

	var OpenIdsString string
	var warnMsgInfoList model.WarnMsgInfoList
	warnMsgInfoList = model.WarnMsgInfoList{
		OpenID:        OpenIdsString,
		SendDate:      util.GetToday(),
		NotDeviceShop: notDeviceShop,
		HaveDiffAbs:   haveDiffShop,
	}
	if err = BODO_Global.BODO_DB.Save(&warnMsgInfoList).Error; err != nil {
		return
	}
	if err = BODO_Global.BODO_DB.Model(model.WarnMsgInfoList{}).Where("send_date = ?", util.GetToday()).First(&warnMsgInfoList).Error; err != nil {
		return
	}
	sendResponse, err := service.ServiceGroupApp.WxMsgService.SendWarningTemplateMsg(wx_request.TemplateMessageRequest{
		ToUser:      OpendIDs,
		Url:         "https://bodo.moonwife.top/#/warnInfoPages?id=" + strconv.Itoa(int(warnMsgInfoList.ID)),
		Type:        "指标类",
		Name:        "销采系统",
		Time:        util.GetToday(),
		Reason:      "存在未安装或偏差较大商户",
		ProjectName: "连云港海州吾悦广场",
	})
	fmt.Println(sendResponse)
	if err != nil {
		return
	}
	for _, v := range sendResponse.OkSendOpenId {
		OpenIdsString += v + ","
	}
	err = BODO_Global.BODO_DB.Model(&warnMsgInfoList).Where("id = ?", warnMsgInfoList.ID).Update("open_id", OpenIdsString).Error
	return
}
