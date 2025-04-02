// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/model/input/apin"
	"hotgo/internal/model/input/frontendIn"

	"github.com/gogf/gf/v2/net/gclient"
)

type (
	IApplication interface {
		EnglishGrammarTranslation(ctx context.Context, req *apin.EnglishGrammarTranslationReq) (res *gclient.Response, err error)
	}
	IFrontendConfig interface {
		GetSmtpSuffix(ctx context.Context) (res *frontendIn.GetSmtpSuffixRes, err error)
	}
	IFrontendMember interface {
		Register(ctx context.Context, in *apin.RegisterReq) (err error)
		Login(ctx context.Context, in *apin.LoginReq) (res *apin.LoginRes, err error)
	}
)

var (
	localApplication    IApplication
	localFrontendConfig IFrontendConfig
	localFrontendMember IFrontendMember
)

func Application() IApplication {
	if localApplication == nil {
		panic("implement not found for interface IApplication, forgot register?")
	}
	return localApplication
}

func RegisterApplication(i IApplication) {
	localApplication = i
}

func FrontendConfig() IFrontendConfig {
	if localFrontendConfig == nil {
		panic("implement not found for interface IFrontendConfig, forgot register?")
	}
	return localFrontendConfig
}

func RegisterFrontendConfig(i IFrontendConfig) {
	localFrontendConfig = i
}

func FrontendMember() IFrontendMember {
	if localFrontendMember == nil {
		panic("implement not found for interface IFrontendMember, forgot register?")
	}
	return localFrontendMember
}

func RegisterFrontendMember(i IFrontendMember) {
	localFrontendMember = i
}
