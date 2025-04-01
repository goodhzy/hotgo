package frontend

import (
	"context"

	"hotgo/internal/model/input/frontendIn"
	"hotgo/internal/service"
)

type sFrontendConfig struct{}

func NewFrontendConfig() *sFrontendConfig {
	return &sFrontendConfig{}
}

func init() {
	service.RegisterFrontendConfig(NewFrontendConfig())
}

func (s *sFrontendConfig) GetSmtpSuffix(ctx context.Context) (res *frontendIn.GetSmtpSuffixRes, err error) {
	smtp, err := service.SysConfig().GetSmtp(ctx)
	if err != nil {
		return nil, err
	}

	res = new(frontendIn.GetSmtpSuffixRes)
	res.Suffix = smtp.Suffix
	return
}
