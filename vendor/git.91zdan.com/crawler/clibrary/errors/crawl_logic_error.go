package errors

import "fmt"

type CrawlLogicError struct {
	err error
}

func NewCrawlAuthorizationError() CrawlLogicError {
	return CrawlLogicError{
		err: fmt.Errorf("用户名或密码错误"),
	}
}

func NewCrawlAuthorizationErrorf(format string, v ...interface{}) CrawlLogicError {
	if v != nil {
		return CrawlLogicError{
			err: fmt.Errorf(format, v),
		}
	}
	return CrawlLogicError{
		err: fmt.Errorf(format),
	}
}


func (e CrawlLogicError) Error() string {
	return e.err.Error()
}
