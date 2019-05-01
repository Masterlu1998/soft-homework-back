package common

import (
	"math"
)

// --------- 封装统一方法 ---------

// GetSuccessObj 成功返回方法
func (resObj *ResObj) GetSuccessObj(code int, prompt string, obj map[string]interface{}) {
	resObj.Code = code
	resObj.Prompt = prompt
	resObj.Obj = obj
}

// GetErrorObj 失败返回方法
func (resObj *ResObj) GetErrorObj(code int, prompt string, err error) {
	resObj.Code = code
	resObj.Prompt = prompt
	resObj.Err = err.Error()
}

// CalculateIncrement 计算增长率并保留两位小数
func CalculateIncrement(a int, b int) float64 {
	increment := math.Round(((float64(b) - float64(a)) / float64(a)) * 100) / 100
	return increment
}
