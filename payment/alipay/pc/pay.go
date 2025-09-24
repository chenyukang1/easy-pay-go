package pc

import (
	"easy-pay-go/payment/alipay"
	"easy-pay-go/pkg/xhttp"
	"encoding/json"
	"reflect"
)

type TradePagePayModel struct {
	OutTradeNo  string `json:"out_trade_no"`
	ProductCode string `json:"product_code"`
	Subject     string `json:"subject"`
	TotalAmount string `json:"total_amount"`
	QrPayMode   string `json:"qr_pay_mode,omitempty"`
	TimeExpire  string `json:"time_expire,omitempty"`
}

func (t TradePagePayModel) ToJson() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type TradePagePayResponse struct {
	Code    string
	Msg     string
	SubCode string
	SubMsg  string
	Body    string
}

func (t *TradePagePayResponse) GetCode() string {
	return t.Code
}

func (t *TradePagePayResponse) GetMsg() string {
	return t.Msg
}

func (t *TradePagePayResponse) GetSubCode() string {
	return t.SubCode
}

func (t *TradePagePayResponse) GetSubMsg() string {
	return t.SubMsg
}

func (t *TradePagePayResponse) GetBody() string {
	return t.Body
}

func (t *TradePagePayResponse) SetCode(code string) {
	t.Code = code
}

func (t *TradePagePayResponse) SetMsg(msg string) {
	t.Msg = msg
}

func (t *TradePagePayResponse) SetSubCode(subCode string) {
	t.SubCode = subCode
}

func (t *TradePagePayResponse) SetSubMsg(subMsg string) {
	t.SubMsg = subMsg
}

func (t *TradePagePayResponse) SetBody(body string) {
	t.Body = body
}

type TradePagePayRequest[T TradePagePayResponse] struct {
	ReturnUrl  string
	NotifyUrl  string
	Encrypt    bool
	BizModel   *TradePagePayModel
	HttpMethod xhttp.HttpMethod
}

func (t *TradePagePayRequest[T]) GetApiMethodName() string {
	return "alipay.trade.page.pay"
}

func (t *TradePagePayRequest[T]) GetApiVersion() string {
	return "1.0"
}

func (t *TradePagePayRequest[T]) GetBizModel() alipay.BizModel {
	return t.BizModel
}

func (t *TradePagePayRequest[T]) GetHttpMethod() xhttp.HttpMethod {
	return t.HttpMethod
}

func (t *TradePagePayRequest[T]) GetNotifyUrl() string {
	return t.NotifyUrl
}

func (t *TradePagePayRequest[T]) GetResponseType() reflect.Type {
	return reflect.TypeOf(&TradePagePayResponse{})
}

func (t *TradePagePayRequest[T]) GetReturnUrl() string {
	return t.ReturnUrl
}

func (t *TradePagePayRequest[T]) HasBizContent() bool {
	return true
}

func (t *TradePagePayRequest[T]) NeedEncrypt() bool {
	return t.Encrypt
}
