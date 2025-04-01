// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package config

import (
	"context"

	"hotgo/api/api/config/v1"
)

type IConfigV1 interface {
	GetSmtpSuffix(ctx context.Context, req *v1.GetSmtpSuffixReq) (res *v1.GetSmtpSuffixRes, err error)
}
