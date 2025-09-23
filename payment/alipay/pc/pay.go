package pc

import (
	"easy-pay-go/payment/alipay"
	"easy-pay-go/pkg/xhttp"
	"encoding/json"
	"reflect"
)

type TradePagePayModel struct {
	Body        string `json:"body"`
	OutTradeNo  string `json:"out_trade_no"`
	ProductCode string `json:"product_code"`
	QrPayMode   string `json:"qr_pay_mode"`
	Subject     string `json:"subject"`
	TimeExpire  string `json:"time_expire"`
	TotalAmount string `json:"total_amount"`
}

func (t TradePagePayModel) ToJson() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type TradePagePayResponse struct {
	Resp *alipay.BaseResponse
}

func (t TradePagePayResponse) SetCode(code string) {
	t.Resp.Code = code
}

func (t TradePagePayResponse) SetMsg(msg string) {
	t.Resp.Msg = msg
}

func (t TradePagePayResponse) SetSubCode(subCode string) {
	t.Resp.SubCode = subCode
}

func (t TradePagePayResponse) SetSubMsg(subMsg string) {
	t.Resp.SubMsg = subMsg
}

func (t TradePagePayResponse) SetBody(body string) {
	t.Resp.Body = body
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
	returnUrl   string
	notifyUrl   string
	needEncrypt bool
	bizModel    *TradePagePayModel
}

func (t TradePagePayRequest[T]) GetApiMethodName() string {
	return "alipay.trade.page.pay"
}

func (t TradePagePayRequest[T]) GetApiVersion() string {
	return "1.0"
}

func (t TradePagePayRequest[T]) GetBizModel() alipay.BizModel {
	return t.bizModel
}

func (t TradePagePayRequest[T]) GetReturnUrl() string {
	return t.returnUrl
}

func (t TradePagePayRequest[T]) GetNotifyUrl() string {
	return t.notifyUrl
}

func (t TradePagePayRequest[T]) GetResponseType() reflect.Type {
	return reflect.TypeOf((*TradePagePayResponse)(nil))
}

func (t TradePagePayRequest[T]) HasBizContent() bool {
	return true
}

func (t TradePagePayRequest[T]) HttpMethod() xhttp.HttpMethod {
	return xhttp.HttpPost
}

func (t TradePagePayRequest[T]) NeedEncrypt() bool {
	return t.needEncrypt
}
