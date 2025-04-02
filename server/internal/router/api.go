// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package router

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/controller/api/application"
	"hotgo/internal/controller/api/common"
	"hotgo/internal/controller/api/config"
	"hotgo/internal/controller/api/member"
	"hotgo/internal/controller/api/pay"

	"hotgo/internal/service"
	"hotgo/utility/simple"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Api 前台路由
func Api(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group(simple.RouterPrefix(ctx, consts.AppApi), func(group *ghttp.RouterGroup) {
		group.Bind(
			pay.NewV1(),    // 支付异步通知
			config.NewV1(), // 配置
			common.NewV1(), // 公共
			application.NewV1(),
		)
		group.Middleware(service.Middleware().ApiAuth)
		group.Bind(
			member.NewV1(), // 管理员
		)
	})
}
