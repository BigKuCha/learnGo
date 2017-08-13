package webcrawler

import "net/http"

//请求
type Request struct {
	httpReq *http.Request //http请求指针
	depth   uint32        //请求深度
}

//创建新的请求
func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq: httpReq, depth: depth}
}

func (req *Request) HttpReq() *http.Request {
	return req.httpReq
}

func (req *Request) Depth() uint32 {
	return req.depth
}

//请求是否有效
func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}
