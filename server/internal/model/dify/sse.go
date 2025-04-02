package dify

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// SSEEvent 表示从服务器接收的单个SSE事件
type SSEEvent struct {
	// 原始数据
	Data string `json:"data"`
	// 解析后的事件类型
	EventType string `json:"event_type"`
}

// SSEWorkflowResponse 工作流SSE响应通用结构
type SSEWorkflowResponse struct {
	// 工作流运行ID
	WorkflowRunID string `json:"workflow_run_id,omitempty"`
	// 任务ID
	TaskID string `json:"task_id,omitempty"`
	// 状态
	Status string `json:"status,omitempty"`
	// 内容块
	Chunk string `json:"chunk,omitempty"`
	// 错误信息
	Error string `json:"error,omitempty"`
	// 完成时的输出
	Outputs map[string]interface{} `json:"outputs,omitempty"`
	// 节点ID
	NodeID string `json:"node_id,omitempty"`
	// 节点类型
	NodeType string `json:"node_type,omitempty"`
	// 消息类型
	MessageType string `json:"message_type,omitempty"`
	// 文件名
	FileName string `json:"file_name,omitempty"`
	// 文件URL
	FileURL string `json:"file_url,omitempty"`
	// 总令牌数
	TotalTokens int `json:"total_tokens,omitempty"`
	// 耗时
	ElapsedTime float64 `json:"elapsed_time,omitempty"`
}

// SSEMessageHandler 处理SSE消息的回调函数类型
type SSEMessageHandler func(event *SSEWorkflowResponse) error

// ParseSSEResponse 解析SSE响应流
func ParseSSEResponse(reader io.Reader, handler SSEMessageHandler) error {
	scanner := bufio.NewScanner(reader)
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
			var response SSEWorkflowResponse
			if err := json.Unmarshal([]byte(data), &response); err != nil {
				return fmt.Errorf("解析SSE事件JSON失败: %w", err)
			}

			// 调用回调函数处理消息
			if err := handler(&response); err != nil {
				return err
			}
		}
	}

	// 检查扫描器错误
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("SSE扫描错误: %w", err)
	}

	return nil
}

// IsComplete 检查事件是否表示流式响应已完成
func (r *SSEWorkflowResponse) IsComplete() bool {
	return r.Status == "succeeded" || r.Status == "failed" || r.Status == "stopped"
}

// IsFirstEvent 检查是否是第一个事件
func (r *SSEWorkflowResponse) IsFirstEvent() bool {
	return r.WorkflowRunID != "" && r.TaskID != ""
}

// ResultText 获取最终结果文本 (适用于文本生成)
func (r *SSEWorkflowResponse) ResultText() string {
	if r.Outputs != nil {
		if result, ok := r.Outputs["result"]; ok {
			if str, ok := result.(string); ok {
				return str
			}
		}
	}
	return ""
}
