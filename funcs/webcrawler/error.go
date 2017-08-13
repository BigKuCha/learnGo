package webcrawler

import (
	"bytes"
	"fmt"
)

type ErrorType string

//爬虫错误接口
type CrawlerError interface {
	Type() ErrorType //获取错误类型
	Error() string   //错误内容
}

//爬虫错误实现
type myCrawlerError struct {
	errType    ErrorType //错误类型
	errMsg     string    //错误提示
	fullErrMsg string    //错误完整提示
}

//条目是否有效
func (item *Item) Valid() bool {
	return item != nil
}

//实现爬虫错误接口方法 Type
func (mce *myCrawlerError) Type() ErrorType {
	return mce.errType
}

//实现爬虫错误接口方法 Error
func (mce *myCrawlerError) Error() string {
	if mce.fullErrMsg == "" {
		mce.genFullErrMsg()
	}
	return mce.fullErrMsg
}

func (mce *myCrawlerError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("Crawler Error:")
	if mce.errType {
		buffer.WriteString(string(mce.errType))
		buffer.WriteString(":")
	}
	buffer.WriteString(mce.errMsg)
	mce.fullErrMsg = fmt.Sprintf("%s\n", mce.fullErrMsg)
	return
}

//创建一个新的爬虫错误
func NewCrawlerError(errType ErrorType, errMsg string) CrawlerError {
	return &myCrawlerError{errType: errType, errMsg: errMsg}
}
