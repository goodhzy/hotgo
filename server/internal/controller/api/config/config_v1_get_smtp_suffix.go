package config

import (
	"context"

	v1 "hotgo/api/api/config/v1"
	"hotgo/internal/service"
)

func (c *ControllerV1) GetSmtpSuffix(ctx context.Context, req *v1.GetSmtpSuffixReq) (res *v1.GetSmtpSuffixRes, err error) {
	smtpSuffixRes, err := service.FrontendConfig().GetSmtpSuffix(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetSmtpSuffixRes{
		List: smtpSuffixRes.Suffix,
	}, nil
}
