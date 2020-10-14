package user_service

import (
	"errors"
	"gf-decoration/app/model/user"
	"gf-decoration/library/base"
	"gf-decoration/library/util"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type LoginRequest struct {
	UserName string `p:"username" v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Password string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
}

type SingUpRequest struct {
	UserName   string `p:"username" v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Password   string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	Repassword string `p:"repassword" v:"required|length:6,30|same:password#请输入密码|密码长度不够|两次密码不一致"`
}

// 通过用户名获取实体
func GetUserByName(username string) (*user.Entity, error) {
	return user.GetUserByName(username)
}

// 通过用户Id获取实体
func GetUserById(id int64) (*user.Entity, error) {
	if id > 0 {
		return user.GetUserById(id)
	}
	return nil, errors.New("没找到该用户")
}

// 用户注册
func SingUp(singupData *SingUpRequest) error {
	salt := util.MD5(util.GetRandomString(8))
	password, err := gmd5.Encrypt(singupData.Password + salt)
	if err != nil {
		return errors.New("系统错误，请稍后再试")
	}
	userEntity := user.Entity{
		UserName: singupData.UserName,
		Password: password,
		Salt:     salt,
	}
	if _, insertErr := user.Model.Insert(userEntity); insertErr != nil {
		return errors.New("注册失败，请稍后再试")
	}
	return nil
}

// 获取缓存的用户信息
func GetCacheUserInfo(r *ghttp.Request) (userInfo *user.Entity) {
	res := base.GfToken.GetTokenData(r)
	gconv.Struct(res.Get("data"), &userInfo)
	return
}

// 获取登录用户 Id
func GetLoginUserId(r *ghttp.Request) (userId int64) {
	userInfo := GetCacheUserInfo(r)
	if userInfo != nil {
		userId = userInfo.Id
	}
	return
}
