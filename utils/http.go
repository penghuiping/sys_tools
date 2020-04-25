package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

//HTTP 用于处理http调用
type HTTP struct {
}

//Get get请求
func (hp *HTTP) Get(url string, headers map[string]string) ([]byte, map[string]string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err1 := client.Do(req)
	if err1 != nil {
		return nil, nil, err1
	}
	defer res.Body.Close()

	headers1 := make(map[string]string)
	for k1, v1 := range res.Header {
		headers1[k1] = strings.Join(v1, ",")
	}

	content, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		return nil, nil, err2
	}
	return content, headers1, nil
}

//Post post请求
func (hp *HTTP) Post(url string, headers map[string]string, body string) ([]byte, map[string]string, error) {
	client := &http.Client{}
	a := strings.NewReader(body)
	req, err := http.NewRequest("POST", url, a)

	if err != nil {
		return nil, nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err1 := client.Do(req)
	if err1 != nil {
		return nil, nil, err1
	}
	defer res.Body.Close()

	headers1 := make(map[string]string)
	for k1, v1 := range res.Header {
		headers1[k1] = strings.Join(v1, ",")
	}

	content, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		return nil, nil, err2
	}
	return content, headers1, nil
}
