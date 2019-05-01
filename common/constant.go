package common

//---------------------- 封装一些统一的类型或结构体

// ResObj : 统一返回结构体
type ResObj struct {
	Code   int    `json:"code"`
	Prompt string `json:"prompt"`
	Obj    Obj    `json:"obj"`
	Err    string  `json:"err"`
}

// Obj : 统一数据格式
type Obj map[string]interface{}

// ResInfo : 统一返回信息和值
type ResInfo struct {
	Code    int
	Message string
}

// 正确返回
var (
	SearchSuccess = ResInfo{Code: 100, Message: "查询成功"}
)

// 错误返回
var (
	SearchFailed    = ResInfo{Code: -101, Message: "查询失败"}
	ParseJSONFailed = ResInfo{Code: -102, Message: "JSON解析失败"}
)
