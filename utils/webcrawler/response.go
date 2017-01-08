package webcrawler

import "net/http"

//响应
type Response struct {
	httpRes *http.Response //响应请求指针
	depth   uint32         //响应深度
}

func NewResponse(httpRes *http.Response, depth uint32) *Response {
	return &Response{httpRes:httpRes, depth: depth}
}

func (res *Response) HttpRes() *http.Response {
	return res.httpRes
}

func (res *Response) Depth() uint32 {
	return res.depth
}

//响应是否有效
func (res *Response) Valid() bool {
	return res.httpRes != nil && res.httpRes.Body != nil
}
