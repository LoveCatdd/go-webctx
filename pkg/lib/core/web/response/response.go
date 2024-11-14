package response

import "github.com/gin-gonic/gin"

type JSON gin.H

// 返回不带error错误， 日志中查看
type Response struct {
	// 状态码
	Code     int
	CodeName string

	// 请求url地址
	Url string

	//传递的消息
	Message string

	// 传递体
	Resp any
}

func json(response Response) JSON {
	return JSON{
		"url":      response.Url,
		"code":     response.Code,
		"codeName": response.CodeName,
		"message":  response.Message,
		"resp":     response.Resp,
	}
}

func SuccessWithResp(resp any) JSON {
	return json(Response{
		Code:     SUCCESS,
		CodeName: StatusName(SUCCESS),
		Resp:     resp,
	})
}

func SuccessWithMessage(message string) JSON {
	return json(Response{
		Code:     SUCCESS,
		CodeName: StatusName(SUCCESS),
		Message:  message,
	})
}

func Success() JSON {
	return json(Response{
		Code:     SUCCESS,
		CodeName: StatusName(SUCCESS),
	})
}

func FailWithMessage(code int, url, message string) JSON {
	return json(Response{
		Code:     code,
		CodeName: StatusName(code),
		Message:  message,
		Url:      url,
	})
}

func Fail(code int, url string) JSON {
	return json(Response{
		Code:     code,
		CodeName: StatusName(code),
		Url:      url,
	})
}

func FailWithResp(code int, url string, resp any) JSON {
	return json(Response{
		Url:      url,
		Code:     code,
		CodeName: StatusName(code),
		Resp:     resp,
	})
}
