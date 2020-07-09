package chttp

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//超时时间
const timeout time.Duration = 5 * time.Second
const DefaultContentType string = "application/json"
const JsonContentType string = DefaultContentType
const XmlContentType string = "application/xml"
const TextContentType string = "text/plain"
const HtmlContentType string = "text/html"
const postMethod string = "POST"
const getMethod string = "GET"

func Post(url string, body string) (string, error) {
	client := &http.Client{Timeout: timeout}
	resp, err := client.Post(url, DefaultContentType, bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println(err.Error())
		return "{}", err
	}
	return rspHttp(resp)
}

func Get(url string) (string, error) {
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return "{}", err
	}
	return rspHttp(resp)
}

func PostHeadBody(url string, body string, head map[string]string) (string, error) {
	req, err := http.NewRequest(postMethod, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println(err.Error())
		return "{}", err
	}
	resp, err := reqHttp(req, head)
	return rspHttp(resp)
}

func PostHeadForm(url string, body map[string]string, head map[string]string) (string, error) {
	req, err := http.NewRequest(postMethod, url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return "{}", err
	}
	for k, v := range body {
		req.PostForm.Set(k, v)
	}
	resp, e := reqHttp(req, head)
	if e != nil {
		return "{}", e
	}
	return rspHttp(resp)
}

func GetHead(url string, head map[string]string) (string, error) {
	req, err := http.NewRequest(getMethod, url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return "{}", err
	}
	resp, e := reqHttp(req, head)
	if e != nil {
		return "{}", e
	}
	return rspHttp(resp)
}

func reqHttp(req *http.Request, head map[string]string) (*http.Response, error) {
	req.Header.Set("Content-Type", DefaultContentType) //设置默认头部
	for k, v := range head {
		req.Header.Set(k, v)
	}
	return (&http.Client{}).Do(req)
}

func rspHttp(rsp *http.Response) (string, error) {
	if rsp == nil {
		return "", errors.New("响应错误!")
	}
	defer rsp.Body.Close()
	result, err := ioutil.ReadAll(rsp.Body)
	return string(result), err
}
