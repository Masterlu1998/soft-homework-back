package common

// 封装一些统一的方法

func (this *ResObj) GetSuccessObj(code int, prompt string, obj map[string]interface{}) {
	this.Code = code
	this.Prompt = prompt
	this.Obj = obj
}

func (this *ResObj) GetErrorObj(code int, prompt string, err error) {
	this.Code = code
	this.Prompt = prompt
	this.Err = err.Error()
}
