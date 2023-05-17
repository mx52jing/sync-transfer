package utils

// 公共返回前端数据
type CommonResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data any `json:"data"`
}

// 接受前端上传数据结构体
type CommonUploadRequest struct {
	Raw string `json:"raw"`
}

const (
	FAILED_CODE = iota - 1 // 普通失败状态码
	SUCCESS_CODE = iota // 成功状态码
)