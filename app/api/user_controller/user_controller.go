package user_controller

import (
	"gf-decoration/app/service/user_service"
	"gf-decoration/library/bean"
	"gf-decoration/library/response"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gvalid"
)

type Controller struct{}

type LoginRequest struct {
	Username string `p:"username" v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Password string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
}

// 用户注册
func (c *Controller) SignUp(r *ghttp.Request) {
	glog.Info("用户注册")
	var singupData *user_service.SingUpRequest
	// 注册信息验证
	if err := r.Parse(&singupData); err != nil {
		if v, ok := err.(*gvalid.Error); ok {
			glog.Error("error when singup valid username and password")
			response.JsonExit(r, 1, v.FirstString(), "")
		}
		response.JsonExit(r, 1, err.Error(), "")
	}
	_, sinErr := user_service.SingUp(singupData)
	if sinErr != nil {
		response.JsonFail(r, sinErr.Error(), "")
	}
	response.JsonSucc(r, "注册成功", "")
}

// 用户登录
func Login(r *ghttp.Request) (string, interface{}) {
	var data *LoginRequest
	if err := r.Parse(&data); err != nil {
		if v, ok := err.(*gvalid.Error); ok {
			glog.Error("error when valid username and password")
			response.JsonExit(r, 1, v.FirstString(), "")
		}
		response.JsonExit(r, 1, err.Error(), "")
	}
	userModel, err := user_service.GetUserByName(data.Username)
	if err != nil {
		response.JsonExit(r, 1, "服务出错，请联系管理员", "")
	}
	if userModel == nil {
		response.JsonExit(r, 1, "没有找到该用户", "")
	}
	glog.Info(userModel.Salt + "---" + data.Password)
	reqPassword, encErr := gmd5.Encrypt(data.Password + userModel.Salt)
	if encErr != nil {
		glog.Error(encErr)
		response.JsonExit(r, 1, "用户名或密码错误")
	}
	if reqPassword != userModel.Password {
		response.JsonExit(r, 1, "用户名或密码错误", "")
	}
	sessionUser := bean.SessionUser{
		Id:       userModel.Id,
		Username: userModel.Username,
		RealName: userModel.RealName,
	}
	return data.Username, sessionUser
}

// 用户登出
func Logout(r *ghttp.Request) bool {
	userId := user_service.GetLoginUserId(r)
	userInfo, err := user_service.GetUserById(userId)
	glog.Info("登出用户信息", userInfo)
	if err != nil {
		response.JsonFail(r, err.Error(), "")
	} else if userInfo.Id != userId {
		response.JsonFail(r, "登出用户不存在", "")
	}
	glog.Info("登出 返回true")
	return true
}

// 获取登录用户信息
func (c *Controller) Info(r *ghttp.Request) {
	id := user_service.GetLoginUserId(r)
	userInfo, err := user_service.GetUserById(id)
	if err != nil {
		response.JsonExit(r, 1, err.Error(), "")
	}
	if userInfo != nil {
		userInfo.Password = ""
		userInfo.Salt = ""
		response.JsonExit(r, 0, "", userInfo)
	}
	response.JsonExit(r, 1, "未找到用户")
}
