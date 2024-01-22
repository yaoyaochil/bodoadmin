package service

import (
	"errors"
	"github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/model/system"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/Wxglobal"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/model"
	systemService "github.com/yaoyaochil/bodo-admin-server/server/service"
	utils2 "github.com/yaoyaochil/bodo-admin-server/server/utils"
	"gorm.io/gorm"
	"time"
)

type WxLoginService struct{}

func (e *WxLoginService) LoginOrCreate(user model.WXUserInfo, id uint) (*system.SysUser, error) {
	if id != 0 {
		//绑定逻辑
		err := global.BODO_DB.First(&model.WXUserInfo{}, "openid = ?", user.Openid).Error
		if err == nil {
			return nil, errors.New("此微信已绑定其他账号")
		}
		user.BODOUserId = id
		err = global.BODO_DB.Create(user).Error
		if err != nil {
			return nil, errors.New("此账号已绑定其他微信")
		}
		return nil, nil
	} else {
		//	扫码登录
		var loginUser model.WXUserInfo
		err := global.BODO_DB.First(&loginUser, "openid = ?", user.Openid).Error
		if err != nil {
			return nil, err
		}
		var resUser system.SysUser
		err = global.BODO_DB.First(&resUser, "id = ?", loginUser.BODOUserId).Error
		if err != nil {
			return nil, err
		}
		return &resUser, nil
	}
}

func (e *WxLoginService) ClearWx(ID uint) error {
	return global.BODO_DB.Delete(&model.WXUserInfo{}, "bodo_user_id = ?", ID).Error
}

var userService = systemService.ServiceGroupApp.SystemServiceGroup.UserService

func (e *WxLoginService) Register(user model.Register, wxUser model.WXUserInfo) (*system.SysUser, error) {
	var regUser system.SysUser
	err := global.BODO_DB.First(&regUser, "username = ?", user.Username).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//	注册
			regUser = system.SysUser{
				Username:    user.Username,
				Password:    user.Password,
				NickName:    "微信注册用户" + wxUser.Openid,
				AuthorityId: Wxglobal.GlobalConfig.AuthorityID,
				Authorities: []system.SysAuthority{
					{
						AuthorityId: Wxglobal.GlobalConfig.AuthorityID,
					},
				},
			}
			bodoInfo, err := userService.Register(regUser)
			if err != nil {
				return nil, err
			}
			wxUser.BODOUserId = bodoInfo.ID
			err = global.BODO_DB.Create(wxUser).Error
			if err != nil {
				return nil, err
			}
			return &bodoInfo, nil
		}
		return nil, err
	} else {
		if ok := utils2.BcryptCheck(user.Password, regUser.Password); !ok {
			return nil, errors.New("密码错误")
		}
		wxUser.BODOUserId = regUser.ID
		err = global.BODO_DB.Create(wxUser).Error
		if err != nil {
			return nil, err
		}
	}
	return &regUser, nil
}

func (e *WxLoginService) GetLoginPic(loginFlag string) (loginPicRes *model.LoginPicRes, err error) {
	pic, err := Wxglobal.GlobalConfig.OfficialAccountApp.QRCode.Temporary(Wxglobal.GlobalConfig.NullCtx, loginFlag, 60*5)
	if err != nil {
		return
	}
	Wxglobal.LoginCodeMap.Store(loginFlag, Wxglobal.CodeInfo{
		CanLogin:   false,
		OpenId:     "",
		CreateTime: time.Now().Unix(),
	})
	return (*model.LoginPicRes)(pic), nil
}

// CheckBindWx 根据用户ID检查是否绑定微信
func (e *WxLoginService) CheckBindWx(ID uint) (bool, error) {
	var wxUser model.WXUserInfo
	err := global.BODO_DB.First(&wxUser, "bodo_user_id = ?", ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
