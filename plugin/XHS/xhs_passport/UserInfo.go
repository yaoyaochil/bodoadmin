package xhs_passport

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/XHS/xhs_request"
	"github.com/yaoyaochil/bodo-admin-server/server/utils"
)

type UserInfoPassport struct{}

func (u *UserInfoPassport) GetUserInfoByUserID(info xhs_request.UserInfoByIDRequest) (data xhs_request.UserInfoByUserIDResponse, err error) {
	url := fmt.Sprintf("https://edith.xiaohongshu.com/api/sns/v3/user/info?profile_page_head_exp=%s&user_id=%s", info.ProfilePageHeadExpansion, info.UserID)
	headers := map[string]string{
		"user-agent":       "discover/8.1 (iPhone; iOS 17.0; Scale/3.00) Resolution/1284*2778 Version/8.1 Build/8010062 Device/(Apple Inc.;iPhone14,3) NetType/WiFi",
		"Host":             "edith.xiaohongshu.com",
		"xy-direction":     "73",
		"accept-language":  "zh-Hans-CN;q=1",
		"shield":           "XYAAAAAQAAAAEAAABTAAAAUzUWEe4xG1IYD9/c+qCLOlKGmTtFa+lG434OfeFWQq8zzPKwlsdhRJ2O/OZfz8Qtrsp+2asyZAw9FBH8ZMr91S8Tju0pgPNz18r5EbI8Hsnu4lpV",
		"xy-platform-info": "platform=iOS&version=8.1&build=8010062&deviceId=A9F08A68-F360-47B7-890D-C5DD0A994C89&bundle=com.xingin.discover",
		"mode":             "gslb",
		"xy-common-params": "app_id=ECFAAF02&build=8010062&channel=AppStore&deviceId=A9F08A68-F360-47B7-890D-C5DD0A994C89&device_fingerprint=202201181017191beb3f5a66a5611fab27a7525fa00cd90103d5d1bdec690c&device_fingerprint1=202201181017191beb3f5a66a5611fab27a7525fa00cd90103d5d1bdec690c&device_model=phone&fid=1642472236-0-0-92106337dfd6d65279e6961acc736d88&gid=7dc4e9a4ce7c556758cd8a8a8ca17e955300e3d847359e2f779098ca&identifier_flag=0&lang=zh-Hans&launch_id=716872553&platform=iOS&project_id=ECFAAF&sid=session.1691906090701657620573&t=1695182144&teenager=0&tz=Asia/Shanghai&uis=light&version=8.1",
		"accept":           "*/*",
	}
	method := "GET"
	response, err := utils.SendRequest(url, method, "", headers)
	if err != nil {
		return xhs_request.UserInfoByUserIDResponse{}, errors.New("请求失败")
	}
	if err = json.Unmarshal(response, &data); err != nil {
		return
	}
	return
}

// GetUserShopItemList 获取商品列表数据 新品倒序
func (u *UserInfoPassport) GetUserShopItemList(info xhs_request.GetUserShopRequest) (data xhs_request.GetUserShopResponse, err error) {
	url := fmt.Sprintf("https://edith.xiaohongshu.com/api/store/personal/get_trade_listing_items?page=%s&seller_id=%s&sort=%s", info.Page, info.SellerId, info.Sort)
	headers := map[string]string{
		"Host":             "edith.xiaohongshu.com",
		"user-agent":       "discover/8.1 (iPhone; iOS 17.0; Scale/3.00) Resolution/1284*2778 Version/8.1 Build/8010062 Device/(Apple Inc.;iPhone14,3) NetType/WiFi",
		"xy-direction":     "73",
		"accept-language":  "zh-Hans-CN;q=1",
		"shield":           "XYAAAAAQAAAAEAAABTAAAAUzUWEe4xG1IYD9/c+qCLOlKGmTtFa+lG434OfeFWQq8zzPKwlsdhRJ2O/OZfz8Qtrsp+2asyZAw9FBH8ZMr91S8Tju1B/wSWxVniTUdul5cnSBh0",
		"xy-platform-info": "platform=iOS&version=8.1&build=8010062&deviceId=A9F08A68-F360-47B7-890D-C5DD0A994C89&bundle=com.xingin.discover",
		"mode":             "gslb",
		"xy-common-params": "app_id=ECFAAF02&build=8010062&channel=AppStore&deviceId=A9F08A68-F360-47B7-890D-C5DD0A994C89&device_fingerprint=202201181017191beb3f5a66a5611fab27a7525fa00cd90103d5d1bdec690c&device_fingerprint1=202201181017191beb3f5a66a5611fab27a7525fa00cd90103d5d1bdec690c&device_model=phone&fid=1642472236-0-0-92106337dfd6d65279e6961acc736d88&gid=7dc4d3bdaa0d556758cd83f08ca17e9553002d08473591177737912e&identifier_flag=0&lang=zh-Hans&launch_id=716916697&platform=iOS&project_id=ECFAAF&sid=session.1691906090701657620573&t=1695225298&teenager=0&tz=Asia/Shanghai&uis=light&version=8.1",
		"accept":           "*/*",
	}
	method := "GET"
	response, err := utils.SendRequest(url, method, "", headers)
	if err != nil {
		return xhs_request.GetUserShopResponse{}, errors.New("请求失败")
	}
	if err = json.Unmarshal(response, &data); err != nil {
		return
	}
	return
}
