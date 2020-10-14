package bean

type SessionUser struct {
	Id       int64  `form:"id" json:"id"`             // 主键
	UserName string `form:"username" json:"username"` // 登录名
}
