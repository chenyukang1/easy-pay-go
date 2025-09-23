package alipay

import (
	"easy-pay-go/pkg/xhttp"
	"reflect"
)

const (
	GatewayPath        = "openapi.alipay.com/gateway.do"
	GatewaySandboxPath = "openapi-sandbox.dl.alipaydev.com/gateway.do"
	TimeFormat         = "2006-01-02 15:04:05"
	UTF8               = "utf-8"
	JSON               = "json"
	RSA2               = "RSA2"
)

type BizModel interface {
	ToJson() (string, error)
}

type Request[T Response] interface {
	GetApiMethodName() string
	GetApiVersion() string
	GetBizModel() BizModel
	GetReturnUrl() string
	GetNotifyUrl() string
	GetResponseType() reflect.Type
	HasBizContent() bool
	HttpMethod() xhttp.HttpMethod
	NeedEncrypt() bool
}

type Response interface {
	GetCode() string
	GetMsg() string
	GetSubCode() string
	GetSubMsg() string
	GetBody() string
	SetCode(code string)
	SetMsg(msg string)
	SetSubCode(subCode string)
	SetSubMsg(subMsg string)
	SetBody(body string)
}

type BaseResponse struct {
	Code    string
	Msg     string
	SubCode string
	SubMsg  string
	Body    string
}

func (response *BaseResponse) GetCode() string {
	return response.Code
}

func (response *BaseResponse) GetMsg() string {
	return response.Msg
}

func (response *BaseResponse) GetSubCode() string {
	return response.SubCode
}

func (response *BaseResponse) GetSubMsg() string {
	return response.SubMsg
}

func (response *BaseResponse) GetBody() string {
	return response.Body
}
