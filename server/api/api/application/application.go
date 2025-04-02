// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package application

import (
	"context"

	"hotgo/api/api/application/v1"
)

type IApplicationV1 interface {
	EnglishGrammarTranslation(ctx context.Context, req *v1.EnglishGrammarTranslationReq) (res *v1.EnglishGrammarTranslationRes, err error)
}
