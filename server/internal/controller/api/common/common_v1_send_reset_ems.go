package common

import (
	"context"

	v1 "hotgo/api/api/common/v1"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

func (c *ControllerV1) SendResetEms(ctx context.Context, req *v1.SendResetEmsReq) (res *v1.SendResetEmsRes, err error) {
	err = service.SysEmsLog().Send(ctx, &sysin.SendEmsInp{
		Event: consts.EmsTemplateResetPwd,
		Email: req.Email,
	})
	return
}
