package member

import (
	"context"

	v1 "hotgo/api/api/member/v1"
	"hotgo/internal/service"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = service.FrontendMember().Register(ctx, &req.RegisterReq)
	return
}
