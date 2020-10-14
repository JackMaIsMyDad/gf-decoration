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
	sinErr := user_service.SingUp(singupData)
	if sinErr != nil {
		response.JsonFail(r, sinErr.Error(), "")
	}
	response.JsonSucc(r, "注册成功", "")
}

// 用户登录
func Login(r *ghttp.Request) (string, interface{}) {
	var data *user_service.LoginRequest
	if err := r.Parse(&data); err != nil {
		if v, ok := err.(*gvalid.Error); ok {
			glog.Error("error when valid username and password")
			response.JsonFail(r, v.FirstString(), "")
		}
		response.JsonFail(r, err.Error(), "")
	}
	userModel, err := user_service.GetUserByName(data.UserName)
	if err != nil {
		response.JsonFail(r, "服务出错，请联系管理员", "")
	}
	if userModel == nil {
		response.JsonFail(r, "没有找到该用户", "")
	}
	glog.Info(userModel.Salt + "---" + data.Password)
	reqPassword, encErr := gmd5.Encrypt(data.Password + userModel.Salt)
	if encErr != nil {
		glog.Error(encErr.Error())
		response.JsonFail(r, "用户名或密码错误", "")
	}
	if reqPassword != userModel.Password {
		response.JsonFail(r, "用户名或密码错误", "")
	}
	sessionUser := bean.SessionUser{
		Id:       userModel.Id,
		UserName: userModel.UserName,
	}
	return data.UserName, sessionUser
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
	return true
}

// 获取登录用户信息
func (c *Controller) Info(r *ghttp.Request) {
	id := user_service.GetLoginUserId(r)
	userInfo, err := user_service.GetUserById(id)
	if err != nil {
		response.JsonFail(r, err.Error(), "")
	}
	if userInfo != nil {
		userInfo.Password = ""
		userInfo.Salt = ""
		response.JsonSucc(r, "获取用户信息成功", userInfo)
	}
	response.JsonFail(r, "未找到用户", "")
}
