package v1

import "github.com/gogf/gf/v2/frame/g"

// SendResetEmsReq 发送重置密码邮件
type SendResetEmsReq struct {
	g.Meta `path:"/ems/sendReset" tags:"邮件" method:"post" summary:"发送重置密码邮件"`
	Email  string `json:"email" v:"required#邮箱不能为空" dc:"邮箱"`
}

type SendResetEmsRes struct {
}

// SendRegisterEmsReq 发送注册邮件
type SendRegisterEmsReq struct {
	g.Meta `path:"/ems/sendRegister" tags:"邮件" method:"post" summary:"发送注册邮件"`
	Email  string `json:"email" v:"required#邮箱不能为空" dc:"邮箱"`
}

type SendRegisterEmsRes struct {
}
