//	路由分发器
//
package main

import (
	"errors"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

type Router struct {
	Service    JZLService
	Processors map[string]JZLProcessor
}

//路由设置
//数据分发
func (this *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//路由分发设置，用来判断url是否合法，通过配置文件的正则表达式配置
	version, resource, token, err := this.ParseURL(r.RequestURI)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
	} else {
		processor, err := this.GetProcessor(resource)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, err.Error())
		} else {
			//处理业务逻辑
			result, err := this.Service.Process(r, version, resource, token, processor)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, err.Error())
			} else {
				io.WriteString(w, result)
			}
		}
	}

	return
}

func (this *Router) GetProcessor(resource string) (JZLProcessor, error) {
	processor, found := this.Processors[resource]
	if found && processor != nil {
		return processor, nil
	} else {
		return nil, errors.New("processor not found.")
	}
}

//
//通过正则表达式选择路由程序
//
func (this *Router) ParseURL(url string) (version int, resource string, token string, err error) {
	//确定是否是本服务能提供的控制类型
	urlPattern, _ := g_config.Get("urlpattern")
	urlRegexp, err := regexp.Compile(urlPattern)
	if err != nil {
		return
	}
	matchs := urlRegexp.FindStringSubmatch(url)
	if matchs == nil {
		err = errors.New("Wrong Request URL")
		return
	}

	versionNum, _ := strconv.ParseInt(matchs[1], 10, 8)
	version = int(versionNum)
	resource = matchs[2]
	token = matchs[3]

	return
}
