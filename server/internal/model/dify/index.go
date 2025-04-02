package dify

// FileInput 文件输入结构
type FileInput struct {
	Type           string `json:"type"`            // document/image/audio/video/custom
	TransferMethod string `json:"transfer_method"` // remote_url/local_file
	URL            string `json:"url,omitempty"`   // 当transfer_method为remote_url时使用
	UploadFileID   string `json:"upload_file_id"`  // 当transfer_method为local_file时使用
}

// WorkflowRunRequest 工作流运行请求
type WorkflowRunRequest struct {
	Inputs       map[string]interface{} `json:"inputs"`            // 输入参数
	ResponseMode string                 `json:"response_mode"`     // streaming/blocking
	User         string                 `json:"user"`              // 用户标识
	Headers      map[string]string      `json:"headers,omitempty"` // 请求头
}

// WorkflowRunData 工作流运行数据
type WorkflowRunData struct {
	ID          string                 `json:"id"`
	WorkflowID  string                 `json:"workflow_id"`
	Status      string                 `json:"status"` // running/succeeded/failed/stopped
	Outputs     map[string]interface{} `json:"outputs,omitempty"`
	Error       string                 `json:"error,omitempty"`
	ElapsedTime float64                `json:"elapsed_time,omitempty"`
	TotalTokens int                    `json:"total_tokens,omitempty"`
	TotalSteps  int                    `json:"total_steps"`
	CreatedAt   int64                  `json:"created_at"`
	FinishedAt  int64                  `json:"finished_at"`
}

// WorkflowRunResponse 工作流运行响应
type WorkflowRunResponse struct {
	WorkflowRunID string          `json:"workflow_run_id"`
	TaskID        string          `json:"task_id"`
	Data          WorkflowRunData `json:"data"`
}

// WorkflowRunDetail 工作流运行详情
type WorkflowRunDetail struct {
	ID          string                 `json:"id"`
	WorkflowID  string                 `json:"workflow_id"`
	Status      string                 `json:"status"`
	Inputs      string                 `json:"inputs"`
	Outputs     map[string]interface{} `json:"outputs"`
	Error       string                 `json:"error"`
	TotalSteps  int                    `json:"total_steps"`
	TotalTokens int                    `json:"total_tokens"`
	CreatedAt   string                 `json:"created_at"`
	FinishedAt  string                 `json:"finished_at"`
	ElapsedTime float64                `json:"elapsed_time"`
}

// FileUploadResponse 文件上传响应
type FileUploadResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Size      int    `json:"size"`
	Extension string `json:"extension"`
	MimeType  string `json:"mime_type"`
	CreatedBy string `json:"created_by"`
	CreatedAt int64  `json:"created_at"`
}

// WorkflowRunInfo 工作流运行信息
type WorkflowRunInfo struct {
	ID          string  `json:"id"`
	Version     string  `json:"version"`
	Status      string  `json:"status"`
	Error       string  `json:"error"`
	ElapsedTime float64 `json:"elapsed_time"`
	TotalTokens int     `json:"total_tokens"`
	TotalSteps  int     `json:"total_steps"`
	CreatedAt   int64   `json:"created_at"`
	FinishedAt  int64   `json:"finished_at"`
}

// EndUserInfo 终端用户信息
type EndUserInfo struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	IsAnonymous bool   `json:"is_anonymous"`
	SessionID   string `json:"session_id"`
}

// WorkflowLog 工作流日志
type WorkflowLog struct {
	ID               string          `json:"id"`
	WorkflowRun      WorkflowRunInfo `json:"workflow_run"`
	CreatedFrom      string          `json:"created_from"`
	CreatedByRole    string          `json:"created_by_role"`
	CreatedByAccount string          `json:"created_by_account"`
	CreatedByEndUser EndUserInfo     `json:"created_by_end_user"`
	CreatedAt        int64           `json:"created_at"`
}

// WorkflowLogsResponse 工作流日志响应
type WorkflowLogsResponse struct {
	Page    int           `json:"page"`
	Limit   int           `json:"limit"`
	Total   int           `json:"total"`
	HasMore bool          `json:"has_more"`
	Data    []WorkflowLog `json:"data"`
}

// AppInfo 应用信息
type AppInfo struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// ImageConfig 图片配置
type ImageConfig struct {
	Enabled         bool     `json:"enabled"`
	NumberLimits    int      `json:"number_limits"`
	Detail          string   `json:"detail"`
	TransferMethods []string `json:"transfer_methods"`
}

// FileUploadConfig 文件上传配置
type FileUploadConfig struct {
	Image ImageConfig `json:"image"`
}

// SystemParameters 系统参数
type SystemParameters struct {
	FileSizeLimit      int `json:"file_size_limit"`
	ImageFileSizeLimit int `json:"image_file_size_limit"`
	AudioFileSizeLimit int `json:"audio_file_size_limit"`
	VideoFileSizeLimit int `json:"video_file_size_limit"`
}

// AppParameters 应用参数
type AppParameters struct {
	UserInputForm    []map[string]interface{} `json:"user_input_form"`
	FileUpload       FileUploadConfig         `json:"file_upload"`
	SystemParameters SystemParameters         `json:"system_parameters"`
}
