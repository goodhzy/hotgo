package apin

type RegisterReq struct {
	Username   string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Code       string `json:"code" v:"required#验证码不能为空" dc:"验证码"`
	Email      string `json:"email" v:"required#邮箱不能为空" dc:"邮箱"`
	Password   string `json:"password" v:"required#密码不能为空" dc:"密码"`
	InviteCode string `json:"invite_code" dc:"邀请码"`
}

type RegisterRes struct {
}
