package httpclient

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"hotgo/internal/model/dify"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
)

// DifyConfig Dify配置
type DifyConfig struct {
	// API基础URL
	BaseURL string
	// API密钥
	ApiKey string
	// 超时时间
	Timeout time.Duration
	// 代理地址
	ProxyURL string
	// 应用ID
	AppID string
}

// DifyClient Dify客户端
type DifyClient struct {
	config DifyConfig
	client *gclient.Client
}

// NewDifyClient 创建Dify客户端
func NewDifyClient(config DifyConfig) (*DifyClient, error) {
	// 设置默认值
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	// 验证必要参数
	if config.BaseURL == "" {
		return nil, fmt.Errorf("baseURL is required")
	}
	if config.ApiKey == "" {
		return nil, fmt.Errorf("apiKey is required")
	}
	if config.AppID == "" {
		return nil, fmt.Errorf("appID is required")
	}

	// 创建客户端
	client := g.Client().Timeout(config.Timeout)

	// 设置代理
	if config.ProxyURL != "" {
		client.SetProxy(config.ProxyURL)
	}

	return &DifyClient{
		config: config,
		client: client,
	}, nil
}

// RunWorkflow 运行工作流
func (c *DifyClient) RunWorkflow(ctx context.Context, req dify.WorkflowRunRequest) (*gclient.Response, error) {
	url := fmt.Sprintf("%s/workflows/run", c.config.BaseURL)

	// 添加应用ID到请求头
	req.Headers = map[string]string{
		"X-App-ID": c.config.AppID,
	}

	resp, err := c.client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.config.ApiKey)).
		SetHeader("Content-Type", "application/json").
		Post(ctx, url, req)
	if err != nil {
		return nil, fmt.Errorf("failed to run workflow: %w", err)
	}

	return resp, nil
}

// StreamHandler 处理流式响应的回调函数类型
type StreamHandler func(chunk string, status string, isDone bool, outputs map[string]interface{}, err error) error

// HandleSSEResponse 处理Dify API的SSE流式响应
func (c *DifyClient) HandleSSEResponse(resp *gclient.Response, handler StreamHandler) error {
	if resp == nil {
		return gerror.New("响应为空")
	}
	defer resp.Body.Close()

	// 检查响应头是否为event-stream
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/event-stream") {
		return gerror.Newf("无效的Content-Type: %s，期望text/event-stream", contentType)
	}

	scanner := bufio.NewScanner(resp.Body)
	var fullText strings.Builder
	var metadata struct {
		WorkflowRunID string
		TaskID        string
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue // 跳过空行
		}

		// 处理data:前缀的行
		if strings.HasPrefix(line, "data:") {
			data := strings.TrimPrefix(line, "data:")
			data = strings.TrimSpace(data)

			// 解析JSON数据
			var event map[string]interface{}
			if err := json.Unmarshal([]byte(data), &event); err != nil {
				return gerror.Newf("解析事件JSON失败: %v", err)
			}

			// 提取关键信息
			status, _ := event["status"].(string)

			// 是否是第一个事件（包含工作流ID和任务ID）
			if wfID, ok := event["workflow_run_id"].(string); ok && wfID != "" {
				metadata.WorkflowRunID = wfID
			}
			if tID, ok := event["task_id"].(string); ok && tID != "" {
				metadata.TaskID = tID
			}

			// 内容块
			if chunk, ok := event["chunk"].(string); ok && chunk != "" {
				fullText.WriteString(chunk)
				// 调用处理器处理内容块
				if err := handler(chunk, status, false, nil, nil); err != nil {
					return err
				}
			}

			// 处理错误
			if errMsg, ok := event["error"].(string); ok && errMsg != "" {
				return handler("", status, true, nil, gerror.New(errMsg))
			}

			// 处理完成
			if status == "succeeded" || status == "failed" || status == "stopped" {
				outputs, _ := event["outputs"].(map[string]interface{})
				return handler(fullText.String(), status, true, outputs, nil)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return gerror.Newf("读取流式响应时出错: %v", err)
	}

	return nil
}

// StreamResponseToHTTP 将Dify SSE响应直接转发到HTTP响应
// 为保持向后兼容性，保留此方法作为对通用函数的封装
func (c *DifyClient) StreamResponseToHTTP(ctx context.Context, resp *gclient.Response, r *ghttp.Request) error {
	return StreamResponseToHTTP(resp, r, SSEStreamOptions())
}

// RunWorkflowBlocking 运行工作流（阻塞模式）并返回解析后的结果
func (c *DifyClient) RunWorkflowBlocking(ctx context.Context, req dify.WorkflowRunRequest) (*dify.WorkflowRunResponse, error) {
	// 确保使用阻塞模式
	req.ResponseMode = "blocking"

	// 发送请求到Dify API
	resp, err := c.RunWorkflow(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析Dify响应
	var difyResp dify.WorkflowRunResponse
	err = json.NewDecoder(resp.Body).Decode(&difyResp)
	if err != nil {
		return nil, gerror.Newf("解析响应失败: %v", err)
	}

	return &difyResp, nil
}
