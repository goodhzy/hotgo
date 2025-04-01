// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package common

import (
	"context"

	"hotgo/api/api/common/v1"
)

type ICommonV1 interface {
	SendResetEms(ctx context.Context, req *v1.SendResetEmsReq) (res *v1.SendResetEmsRes, err error)
	SendRegisterEms(ctx context.Context, req *v1.SendRegisterEmsReq) (res *v1.SendRegisterEmsRes, err error)
}
