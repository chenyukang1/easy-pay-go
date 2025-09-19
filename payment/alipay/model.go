package alipay

const (
	GatewayUrl        = "https://openapi.alipay.com/gateway.do"
	GatewaySandboxUrl = "https://openapi-sandbox.dl.alipaydev.com/gateway.do"
	UTF8              = "utf-8"
	RSA2              = "RSA2"
)

type Request[T Response] interface {
	GetApiMethodName() string
	GetApiVersion() string
	GetReturnUrl() string
	GetNotifyUrl() string
}

type Response interface {
	GetCode() string
	GetMsg() string
	GetSubCode() string
	GetSubMsg() string
	GetBody() string
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
