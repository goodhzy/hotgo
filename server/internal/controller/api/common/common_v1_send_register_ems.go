package common

import (
	"context"

	v1 "hotgo/api/api/common/v1"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

func (c *ControllerV1) SendRegisterEms(ctx context.Context, req *v1.SendRegisterEmsReq) (res *v1.SendRegisterEmsRes, err error) {
	err = service.SysEmsLog().Send(ctx, &sysin.SendEmsInp{
		Event: consts.EmsTemplateRegister,
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}
	return
}
