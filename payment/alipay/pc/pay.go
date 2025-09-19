package pc

import (
	"easy-pay-go/payment/alipay"
)

type TradePagePayResponse struct {
	Resp *alipay.BaseResponse
}

func (t TradePagePayResponse) GetCode() string {
	return t.Resp.Code
}

func (t TradePagePayResponse) GetMsg() string {
	return t.Resp.Msg
}

func (t TradePagePayResponse) GetSubCode() string {
	return t.Resp.SubCode
}

func (t TradePagePayResponse) GetSubMsg() string {
	return t.Resp.SubMsg
}

func (t TradePagePayResponse) GetBody() string {
	return t.Resp.Body
}

type TradePagePayRequest[T TradePagePayResponse] struct {
	apiMethodName string
	apiVersion    string
	returnUrl     string
	notifyUrl     string
}

func (t TradePagePayRequest[T]) GetApiMethodName() string {
	return t.apiMethodName
}

func (t TradePagePayRequest[T]) GetApiVersion() string {
	return t.apiVersion
}

func (t TradePagePayRequest[T]) GetReturnUrl() string {
	return t.returnUrl
}

func (t TradePagePayRequest[T]) GetNotifyUrl() string {
	return t.notifyUrl
}
