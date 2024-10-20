package commonRes

import "fmt"

/*
通用错误处理
*/

type Response struct {

	//code是自定义的code
	Code int `json:"code"`

	//错误的信息，我们口语化的错误，参数错误
	Message string `json:"message"`

	//一些需要返回给用户的错误，比如模板解析出问题，哪里出问题
	//注意，不要乱传递reason，比如gorm执行sql出错，你不能放reason里，不然会暴露sql了
	Reason string `json:"reason"`

	//返回的数据，在http里可以这样做，但是在grpc里不能这么做
	Data interface{} `json:"data"`
}

// 实现error接口，返回给error
func (receiver *Response) Error() string {
	return fmt.Sprintf("code=%d ", receiver.Code) +
		fmt.Sprintf("message=%s ", receiver.Message) +
		fmt.Sprintf("reason=%s ", receiver.Reason)
}

func (receiver *Response) WithReason(reason string) *Response {
	response := NewResponse(receiver.Code, receiver.Message)
	response.Reason = reason
	return response
}

func (receiver *Response) WithData(data interface{}) *Response {
	response := NewResponse(receiver.Code, receiver.Message)
	response.Data = data
	return response
}

func NewResponse(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}
