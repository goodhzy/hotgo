// Package member
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package v1

import (
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/apin"

	"github.com/gogf/gf/v2/frame/g"
)

// GetIdByCodeReq 通过邀请码获取用户ID
type GetIdByCodeReq struct {
	g.Meta `path:"/member/getIdByCode" method:"post" tags:"用户" summary:"通过邀请码获取用户ID"`
	Code   string `json:"code"   dc:"邀请码"`
}

type GetIdByCodeRes struct{}

// RegisterReq 注册用户
type RegisterReq struct {
	g.Meta `path:"/member/register" method:"post" tags:"用户" summary:"注册用户"`
	apin.RegisterReq
}

type RegisterRes struct {
	*adminin.LoginModel
}
