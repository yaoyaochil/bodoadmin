package system

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/model/system"
	"github.com/yaoyaochil/bodo-admin-server/server/model/system/request"
	"github.com/yaoyaochil/bodo-admin-server/server/utils"
	"gorm.io/gorm"
)

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: Register
//@description: 用户注册
//@param: u Model.SysUser
//@return: userInter system.SysUser, err error

type UserService struct{}

func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.BODO_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.BODO_DB.Create(&u).Error
	return u, err
}

// @author: [wangrui19970405](https://github.com/wangrui19970405)
// @author: [wangrui19970405](https://github.com/wangrui19970405)
// @function: Login
// @description: 用户登录
// @param: u *Model.SysUser
// @return: err error, userInter *Model.SysUser
func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.BODO_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = global.BODO_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user)
		// 更新登录时间 登录ip
		err = global.BODO_DB.Model(&user).Updates(system.SysUser{LastLoginTime: u.LastLoginTime, LastLoginIp: u.LastLoginIp}).Error
	}
	return &user, err
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *Model.SysUser, newPassword string
//@return: userInter *Model.SysUser,err error

func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (userInter *system.SysUser, err error) {
	var user system.SysUser
	if err = global.BODO_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.BODO_DB.Save(&user).Error
	return &user, err

}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *UserService) GetUserInfoList(info request.SystemUserRequest) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.BODO_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if info.Phone != "" {
		db = db.Where("phone like ?", "%"+info.Phone+"%")
	}
	if info.NickName != "" {
		db = db.Where("nick_name like ?", "%"+info.NickName+"%")
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func (userService *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {
	assignErr := global.BODO_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.BODO_DB.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error

func (userService *UserService) SetUserAuthorities(id uint, authorityIds []uint) (err error) {
	return global.BODO_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []system.SysUserAuthority
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, system.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func (userService *UserService) DeleteUser(id int) (err error) {
	var user system.SysUser
	err = global.BODO_DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = global.BODO_DB.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
	return err
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser Model.SysUser
//@return: err error, user Model.SysUser

func (userService *UserService) SetUserInfo(req system.SysUser) error {
	return global.BODO_DB.Updates(&req).Error
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.SysUser

func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
	var reqUser system.SysUser
	err = global.BODO_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return reqUser, err
	}
	MenuServiceApp.UserAuthorityDefaultRouter(&reqUser)
	return reqUser, err
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *Model.SysUser

func (userService *UserService) FindUserById(id int) (user *system.SysUser, err error) {
	var u system.SysUser
	err = global.BODO_DB.Where("`id` = ?", id).First(&u).Error
	return &u, err
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *Model.SysUser

func (userService *UserService) FindUserByUuid(uuid string) (user *system.SysUser, err error) {
	var u system.SysUser
	if err = global.BODO_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: resetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error

func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = global.BODO_DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}

// GetUserTotal 获取总用户数
func (userService *UserService) GetUserTotal() (count int64, err error) {
	err = global.BODO_DB.Model(&system.SysUser{}).Count(&count).Error
	return
}
