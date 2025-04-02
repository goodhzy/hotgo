package member

import (
	"context"

	v1 "hotgo/api/api/member/v1"
	"hotgo/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	loginRes, err := service.FrontendMember().Login(ctx, &req.LoginReq)
	if err != nil {
		return nil, err
	}
	return &v1.LoginRes{
		LoginRes: loginRes,
	}, nil
}
