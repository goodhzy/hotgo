package httpclient

import (
	"context"
	"fmt"
	"io"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
)

// StreamOptions 流式转发选项
type StreamOptions struct {
	// Content-Type 头部值
	ContentType string
	// 其他要添加的头部
	AdditionalHeaders map[string]string
	// 读取缓冲区大小
	BufferSize int
}

// SSEStreamOptions 返回用于SSE的默认流选项
func SSEStreamOptions() *StreamOptions {
	return &StreamOptions{
		ContentType: "text/event-stream",
		AdditionalHeaders: map[string]string{
			"Cache-Control":               "no-cache",
			"Connection":                  "keep-alive",
			"Access-Control-Allow-Origin": "*",
		},
		BufferSize: 4096,
	}
}

// JSONStreamOptions 返回用于流式JSON的默认流选项
func JSONStreamOptions() *StreamOptions {
	return &StreamOptions{
		ContentType: "application/x-ndjson",
		AdditionalHeaders: map[string]string{
			"Cache-Control": "no-cache",
			"Connection":    "keep-alive",
		},
		BufferSize: 4096,
	}
}

// BinaryStreamOptions 返回用于二进制流的默认流选项
func BinaryStreamOptions() *StreamOptions {
	return &StreamOptions{
		ContentType: "application/octet-stream",
		AdditionalHeaders: map[string]string{
			"Cache-Control": "no-cache",
		},
		BufferSize: 8192,
	}
}

// StreamResponseToHTTP 将流式响应直接转发到HTTP响应
// 这是一个通用函数，可用于任何类型的流式响应转发到HTTP响应
// 参数:
// - resp: 包含流式内容的响应
// - r: GoFrame请求对象，用于写入响应
// - options: 可选配置，例如自定义头部 (可为nil使用默认值)
func StreamResponseToHTTP(resp *gclient.Response, r *ghttp.Request, options ...*StreamOptions) error {
	// 获取选项，如果未提供则使用默认值
	var opts StreamOptions
	if len(options) > 0 && options[0] != nil {
		opts = *options[0]
	} else {
		// 默认选项 - 适用于SSE
		opts = StreamOptions{
			ContentType: "text/event-stream",
			AdditionalHeaders: map[string]string{
				"Cache-Control":               "no-cache",
				"Connection":                  "keep-alive",
				"Access-Control-Allow-Origin": "*",
			},
			BufferSize: 4096,
		}
	}

	// 设置Content-Type
	r.Response.Header().Set("Content-Type", opts.ContentType)

	// 设置其他头部
	for k, v := range opts.AdditionalHeaders {
		r.Response.Header().Set(k, v)
	}

	defer resp.Body.Close()

	// 直接将流式响应转发给客户端
	buffer := make([]byte, opts.BufferSize)
	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			r.Response.Write(buffer[:n])
			r.Response.Flush()
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			// 发送错误事件
			if opts.ContentType == "text/event-stream" {
				// 对于SSE，发送格式化的错误事件
				errEvent := fmt.Sprintf("data: {\"error\": \"%s\", \"status\": \"failed\"}\n\n", err.Error())
				r.Response.WriteString(errEvent)
			} else {
				// 对于其他类型，仅记录错误
				g.Log().Errorf(context.Background(), "流转发时发生错误: %v", err)
			}
			r.Response.Flush()
			return err
		}
	}

	return nil
}
