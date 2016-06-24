package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

type JZLProcessor interface {
	ProcessFunc(version int, resource string, method string, token string, params map[string]string, body []byte, result map[string]interface{}) error
}
type JZLService interface {
	Process(r *http.Request, version int, resource string, token string, processor JZLProcessor) (string, error)
}

type BaseService struct {
	RequestParams map[string]string
}

func (this *BaseService) Process(r *http.Request, version int, resource string, token string, processor JZLProcessor) (string, error) {
	var err error
	var body []byte

	result := make(map[string]interface{})

	//解析参数
	err = this.parseArgs(r)
	if err != nil {
		result["error_code"] = -1
		result["message"] = err.Error()
		goto END
	}

	body, err = ioutil.ReadAll(r.Body)
	if err != nil && err != io.EOF {
		result["error_code"] = -1
		result["message"] = err.Error()
		goto END
	}

	//处理逻辑
	err = processor.ProcessFunc(version, resource, r.Method, token, this.RequestParams, body, result)
	if err != nil {
		result["error_code"] = -1
		result["message"] = err.Error()
		goto END
	}

	result["error_code"] = 0

END:
	resStr, _ := this.buildResult(result)
	return resStr, err
}

func (this *BaseService) buildResult(result map[string]interface{}) (string, error) {
	echo, _ := result["echo"].(string)
	return echo, nil
}

func (this *BaseService) parseArgs(r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	//每次都重新生成一个新的map，否则之前请求的参数会保留其中
	this.RequestParams = make(map[string]string)
	for k, v := range r.Form {
		this.RequestParams[k] = v[0]
	}

	return nil
}
