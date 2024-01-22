package model

import "github.com/yaoyaochil/bodo-admin-server/server/global"

// UserSummary 用户增减数据
type UserSummary struct {
	global.BODO_MODEL
	RefDate      string `json:"ref_date"`      // 数据的日期，需在begin_date和end_date之间
	UserSource   int    `json:"user_source"`   // 用户的渠道，数值代表的含义如下： 0代表其他合计， 1代表公众号搜索， 17代表名片分享， 30代表扫描二维码， 51代表支付后关注（在支付完成页）， 57代表文章内账号名称 100代表微信广告， 161代表他人转载， 149代表小程序关注， 200代表视频号， 201代表直播
	NewUser      int    `json:"new_user"`      // 新增的用户数量
	CancelUser   int    `json:"cancel_user"`   // 取消关注的用户数量，new_user减去cancel_user即为净增用户数量
	CumulateUser int    `json:"cumulate_user"` // 总用户量
}
