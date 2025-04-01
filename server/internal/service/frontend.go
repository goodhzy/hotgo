// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/model/input/frontendIn"
)

type (
	IFrontendConfig interface {
		GetSmtpSuffix(ctx context.Context) (res *frontendIn.GetSmtpSuffixRes, err error)
	}
)

var (
	localFrontendConfig IFrontendConfig
)

func FrontendConfig() IFrontendConfig {
	if localFrontendConfig == nil {
		panic("implement not found for interface IFrontendConfig, forgot register?")
	}
	return localFrontendConfig
}

func RegisterFrontendConfig(i IFrontendConfig) {
	localFrontendConfig = i
}
