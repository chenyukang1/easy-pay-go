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
	JSON               = "JSON"
	RSA2               = "RSA2"
)

type BizModel interface {
	ToJson() (string, error)
}

type Request[T Response] interface {
	GetApiMethodName() string
	GetApiVersion() string
	GetBizModel() BizModel
	GetHttpMethod() xhttp.HttpMethod
	GetNotifyUrl() string
	GetReturnUrl() string
	GetResponseType() reflect.Type
	HasBizContent() bool
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
