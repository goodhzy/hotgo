package frontend

import (
	"context"
	"hotgo/internal/httpclient"
	"hotgo/internal/model/dify"
	"hotgo/internal/model/input/apin"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/net/gclient"
)

type sApplication struct{}

func NewApplication() *sApplication {
	return &sApplication{}
}

func init() {
	service.RegisterApplication(NewApplication())
}

func (l *sApplication) EnglishGrammarTranslation(ctx context.Context, req *apin.EnglishGrammarTranslationReq) (res *gclient.Response, err error) {
	// 创建Dify客户端
	client, err := httpclient.NewDifyClient(httpclient.DifyConfig{
		BaseURL:  "http://119.29.157.38:81/v1",
		ApiKey:   "app-tm9VCda3thkRafWUbm70cyOL",
		AppID:    "39lrexqkyz5j008i",
		ProxyURL: "192.168.1.167:8888",
	})
	if err != nil {
		return nil, err
	}

	// 执行工作流获取响应
	return client.RunWorkflow(ctx, dify.WorkflowRunRequest{
		Inputs: map[string]interface{}{
			"text": req.Text,
		},
		ResponseMode: "streaming",
		User:         "test",
	})

}
