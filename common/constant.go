package common

type ResObj struct {
	Code   int                    `json:"code"`
	Prompt string                 `json:"prompt"`
	Obj    map[string]interface{} `json:"obj"`
	Err    error                  `json:"err"`
}

func (this *ResObj) GetSuccessObj(code int, prompt string, obj map[string]interface{}) {
	this.Code = code
	this.Prompt = prompt
	this.Obj = obj
}

func (this *ResObj) GetErrorObj(code int, prompt string, err error) {
	this.Code = code
	this.Prompt = prompt
	this.Err = err
}
