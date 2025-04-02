package application

import (
	"context"

	v1 "hotgo/api/api/application/v1"
	"hotgo/internal/httpclient"

	"hotgo/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// EnglishGrammarTranslation 英语语法翻译
func (c *ControllerV1) EnglishGrammarTranslation(ctx context.Context, req *v1.EnglishGrammarTranslationReq) (res *v1.EnglishGrammarTranslationRes, err error) {
	resp, err := service.Application().EnglishGrammarTranslation(ctx, &req.EnglishGrammarTranslationReq)
	if err != nil {
		return nil, err
	}
	// 从上下文获取请求对象
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return nil, gerror.New("无法获取请求对象")
	}

	options := httpclient.SSEStreamOptions()
	err = httpclient.StreamResponseToHTTP(resp, r, options)
	if err != nil {
		return nil, err
	}

	// 由于是流式响应，返回空结果
	return &v1.EnglishGrammarTranslationRes{}, nil
}
