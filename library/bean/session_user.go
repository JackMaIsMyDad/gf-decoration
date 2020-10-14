package bean

type SessionUser struct {
	Id       int64  `form:"id" json:"id"`               // 主键
	Username string `form:"username" json:"username"`   // 登录名
	RealName string `form:"real_name" json:"real_name"` // 真实姓名
}
