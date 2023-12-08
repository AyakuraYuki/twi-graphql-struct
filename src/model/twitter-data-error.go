package model

type ErrorRsp struct {
	Message    string           `json:"message"`     // 错误信息
	Locations  []*ErrorLocation `json:"locations"`   // 错误内容在原始文案的位置信息
	Path       []any            `json:"path"`        // 错误数据节点的路径
	Extensions *ErrorExtensions `json:"extensions"`  // 错误信息的扩展内容
	Code       int              `json:"code"`        // 错误代码
	Kind       string           `json:"kind"`        // 错误类型
	Name       string           `json:"name"`        // 错误名称
	Source     string           `json:"source"`      // 错误来源
	RetryAfter int              `json:"retry_after"` // 重试次数
	Tracing    *ErrorTracing    `json:"tracing"`     // 链路追踪信息
}

type ErrorLocation struct {
	Line   int `json:"line"`   // 格式化JSON所在行号
	Column int `json:"column"` // 格式化JSON所在列号
}

type ErrorExtensions struct {
	Name       string        `json:"name"`        // 错误名称
	Source     string        `json:"source"`      // 错误来源
	RetryAfter int           `json:"retry_after"` // 重试次数
	Code       int           `json:"code"`        // 错误代码
	Kind       string        `json:"kind"`        // 错误类型
	Tracing    *ErrorTracing `json:"tracing"`     // 链路追踪信息
}

type ErrorTracing struct {
	TraceId string `json:"trace_id"` // trace_id
}
